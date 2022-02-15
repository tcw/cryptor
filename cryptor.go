package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha512"
	b64 "encoding/base64"
	"errors"
	"flag"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	encrypt := flag.Bool("e", false, "encrypt")
	decrypt := flag.Bool("d", false, "decrypt")

	flag.Usage = func() {
		_, err := fmt.Fprintf(os.Stderr, "Usage: %s [options] [source] [desination]\n", os.Args[0])
		if err != nil {
			return
		}
		flag.PrintDefaults()
	}
	flag.Parse()
	if *encrypt == *decrypt {
		fmt.Println("chose either option -e or -d")
		return
	}

	args := flag.Args()
	if len(args) != 2 {
		fmt.Println("missing file as input parameter, should be [options] [source] [desination]")
		return
	}
	sourceFile := args[0]
	destinationFile := args[1]
	sourcePath, err := filepath.Abs(sourceFile)
	check(err)
	destinationPath, err := filepath.Abs(destinationFile)
	check(err)
	source, err := os.OpenFile(sourcePath, os.O_RDONLY, 0660)
	check(err)
	defer source.Close()
	destination, err := os.OpenFile(destinationPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0660)
	defer destination.Close()
	check(err)

	if *encrypt {
		key := secret(true)
		fileContent, err := io.ReadAll(source)
		check(err)
		encoded, err := encryptAES(key, string(fileContent))
		check(err)
		_, err = destination.Write([]byte(encoded))
		check(err)
		return
	}
	if *decrypt {
		key := secret(false)
		fileContent, err := io.ReadAll(source)
		check(err)
		decoded, err := decryptAES(key, string(fileContent))
		_, err = destination.Write([]byte(decoded))
		check(err)
		return
	}
}

func secret(reenter bool) []byte {
	fmt.Print("enter key: ")
	bytePassword, err := terminal.ReadPassword(0)
	fmt.Println()
	check(err)
	if reenter {
		fmt.Printf("re-enter key: ")
		retypedBytePassword, err := terminal.ReadPassword(0)
		check(err)
		fmt.Println()
		if string(bytePassword) != string(retypedBytePassword) {
			log.Fatalln("keys have to match")
		}
	}
	password := string(bytePassword)
	trimmedPassword := strings.TrimSpace(password)
	hash := sha512.New512_256()
	hash.Write([]byte(trimmedPassword))
	return hash.Sum(nil)
}

func encryptAES(key []byte, message string) (encoded string, err error) {
	plainText := []byte(message)
	block, err := aes.NewCipher(key)
	check(err)
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		check(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	return b64.RawStdEncoding.EncodeToString(cipherText), err
}
func decryptAES(key []byte, secure string) (decoded string, err error) {
	cipherText, err := b64.RawStdEncoding.DecodeString(secure)
	check(err)
	block, err := aes.NewCipher(key)
	check(err)
	if len(cipherText) < aes.BlockSize {
		err = errors.New("ciphertext block size is too short")
		check(err)
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), err
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

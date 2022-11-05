# Cryptor

A simple commandline tool for encrypting and decrypting file with AES-256 Cipher feedback (CFB).
Both input and output is base64, which means secrets are easy to store in all forms of storage.

## Install

```shell
 go install github.com/tcw/cryptor@latest
```

## Usage

To show usage in commandline
```shell
cryptor --help
```

An example to illustrate basic usage
```shell
echo "This is super secret" > secret.txt
cryptor -e secret.txt secret.aes
> enter key:
> re-enter key:

cryptor -d secret.aes secret2.txt
> enter key:
> re-enter key:

cat secret2.txt
diff secret.txt secret2.txt
```
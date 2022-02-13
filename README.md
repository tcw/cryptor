# Cryptor

A simple commandline tool for encrypting and decrypting file with AES-256 Cipher feedback (CFB).


## Usage 

```shell
echo "This is super secret" > secret.txt
cryptor -e secret.txt secret.aes
```

```shell
cryptor -d secret.aes secret2.txt
```
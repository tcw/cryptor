# Cryptor

A simple commandline tool for encrypting and decrypting file with AES-256 Cipher feedback (CFB).


## Usage 

```shell
go get -U github.com/tcw/cryptor
```

```shell
echo "This is super secret" > secret.txt
cryptor -e secret.txt secret.aes
> enter key:
> re-enter key:

cryptor -d secret.aes secret2.txt
> enter key:
> re-enter key:

diff secret.txt secret2.txt
```
# Cryptor

A simple commandline tool for encrypting and decrypting file with AES-256 Cipher feedback (CFB).


## Install

```shell
 go get -u github.com/tcw/cryptor
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

diff secret.txt secret2.txt
```
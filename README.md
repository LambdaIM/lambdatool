# lambdatool

Command line tool of Lambda Blockchain.

## Features

1. AES encryption and decryption
2. Conversion between `lambvaloper` and `0x` prefix address
3. Calculate the MD5 value of file
4. Supports 64-bit Linux, MacOS, Windows by default

## Compile (x86_64)

### Linux

```shell
GOOS=linux GOARCH=amd64 go build -ldflags '-s -w' -o lambdatool
```

### Windows

```shell
GOOS=windows GOARCH=amd64 go build -ldflags '-s -w' -o lambdatool.exe
```

### MacOS

```shell
GOOS=darwin GOARCH=amd64 go build -ldflags '-s -w' -o lambdatool
```
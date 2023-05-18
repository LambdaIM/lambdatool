package commands

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"github.com/mkideal/cli"
)

type cryptoOpts struct {
	cli.Helper2
	Data   string `cli:"*data" usage:"Raw data that to be encrypted or decrypted"`
	Secret string `cli:"*secret" usage:"Secret used for encryption or decryption, length must be 16 or 24 or 32"`
}

var (
	CmdEncrypt = &cli.Command{
		Name: "encrypt",
		Desc: "Encrypt data using AES and output Base64 encoded result",
		Argv: func() interface{} {
			return new(cryptoOpts)
		},
		Fn: func(ctx *cli.Context) error {
			argv := ctx.Argv().(*cryptoOpts)
			if 0 == len(argv.Data) {
				return errors.New("raw data length is 0")
			}
			result, err := encrypt(argv)
			if nil != err {
				return err
			}
			ctx.String("%s\n", result)
			return nil
		},
	}
	CmdDecrypt = &cli.Command{
		Name: "decrypt",
		Desc: "Decrypt AES-encrypted and Base64-encoded data",
		Argv: func() interface{} {
			return new(cryptoOpts)
		},
		Fn: func(ctx *cli.Context) error {
			argv := ctx.Argv().(*cryptoOpts)
			if 0 == len(argv.Data) {
				return errors.New("raw data length is 0")
			}
			result, err := decrypt(argv)
			if nil != err {
				return err
			}
			ctx.String("%s\n", result)
			return nil
		},
	}
)

func encrypt(opts *cryptoOpts) (string, error) {
	key := []byte(opts.Secret)
	block, err := aes.NewCipher(key)
	if nil != err {
		return "", err
	}
	size := block.BlockSize()
	data := padding([]byte(opts.Data), size)
	mode := cipher.NewCBCEncrypter(block, key[:size])
	encrypted := make([]byte, len(data))
	mode.CryptBlocks(encrypted, data)
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func padding(rawData []byte, size int) []byte {
	content := size - len(rawData)%size
	text := bytes.Repeat([]byte{byte(content)}, content)
	return append(rawData, text...)
}

func decrypt(opts *cryptoOpts) (string, error) {
	encrypted, err := base64.StdEncoding.DecodeString(opts.Data)
	if nil != err {
		return "", err
	}
	key := []byte(opts.Secret)
	block, err := aes.NewCipher(key)
	if nil != err {
		return "", err
	}
	size := block.BlockSize()
	mode := cipher.NewCBCDecrypter(block, key[:size])
	data := make([]byte, len(encrypted))
	mode.CryptBlocks(data, encrypted)
	data = unPadding(data)
	return string(data), nil
}

func unPadding(rawData []byte) []byte {
	length := len(rawData)
	end := int(rawData[length-1])
	return rawData[:(length - end)]
}

package commands

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/mkideal/cli"
	"os"
)

type md5Opts struct {
	cli.Helper2
	File string `cli:"*file" usage:"The target file"`
}

var CmdMD5 = &cli.Command{
	Name: "md5",
	Desc: "Calculate the MD5 value of file",
	Argv: func() interface{} {
		return new(md5Opts)
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*md5Opts)
		if 0 == len(argv.File) {
			return errors.New("file not found")
		}
		value, err := calc(argv.File)
		if nil != err {
			return err
		}
		ctx.String("%s\n", value)
		return nil
	},
}

func calc(file string) (string, error) {
	data, err := os.ReadFile(file)
	if nil != err {
		return "", err
	}
	hash := md5.New()
	hash.Write(data)
	value := hash.Sum(nil)
	return hex.EncodeToString(value), nil
}

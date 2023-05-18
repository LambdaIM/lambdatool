package main

import (
	"fmt"
	"github.com/mkideal/cli"
	"lambdatool/commands"
	"os"
)

var (
	cmdRoot = &cli.Command{
		Desc: "Command line tool of Lambda Blockchain.",
		Argv: func() interface{} {
			return new(cli.Helper2)
		},
		Fn: func(ctx *cli.Context) error {
			ctx.String(ctx.Usage())
			return nil
		},
	}
)

func main() {
	err := cli.Root(cmdRoot,
		cli.Tree(commands.CmdEncrypt),
		cli.Tree(commands.CmdDecrypt),
		cli.Tree(commands.CmdConvert),
		cli.Tree(commands.CmdMD5),
	).Run(os.Args[1:])
	if nil != err {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

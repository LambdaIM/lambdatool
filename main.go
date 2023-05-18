package main

import (
	"fmt"
	"github.com/mkideal/cli"
	"lambdatool/commands"
	"os"
)

const CurrentVersion = "v1.0.1"

type mainOpts struct {
	cli.Helper2
	Version bool `cli:"!v,version" usage:"Version of the command line tool"`
}

var (
	cmdRoot = &cli.Command{
		Desc: "Command line tool of Lambda Blockchain.",
		Argv: func() interface{} {
			return new(mainOpts)
		},
		Fn: func(ctx *cli.Context) error {
			argv := ctx.Argv().(*mainOpts)
			if argv.Version {
				ctx.String("%s\n", CurrentVersion)
			} else {
				ctx.String(ctx.Usage())
			}
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

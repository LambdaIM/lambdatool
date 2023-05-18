package commands

import (
	"errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/mkideal/cli"
	"strings"
)

type addressOpts struct {
	cli.Helper2
	Addr string `cli:"*addr" usage:"Address starting with 'lambvaloper' and '0x'"`
}

var CmdConvert = &cli.Command{
	Name: "address",
	Desc: "Format conversion between addresses starting with 'lambvaloper' and '0x'",
	Argv: func() interface{} {
		return new(addressOpts)
	},
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*addressOpts)
		if 0 == len(argv.Addr) {
			return errors.New("address length is 0")
		}
		result, err := convert(argv.Addr)
		if nil != err {
			return err
		}
		ctx.String("%s\n", result)
		return nil
	},
}

func convert(address string) (string, error) {
	if strings.HasPrefix(address, "0x") {
		return toLambda(address)
	} else if strings.HasPrefix(address, "lambvaloper") {
		return to0x(address)
	}
	return "", errors.New("unknown address prefix")
}

func to0x(address string) (string, error) {
	valAddr, err := sdk.ValAddressFromBech32(address)
	if nil != err {
		return "", err
	}
	return common.BytesToAddress(valAddr).Hex(), nil
}

func toLambda(address string) (string, error) {
	accAddr, err := sdk.Bech32ifyAddressBytes(sdk.GetConfig().GetBech32ValidatorAddrPrefix(), common.HexToAddress(address).Bytes())
	if nil != err {
		return "", err
	}
	return accAddr, nil
}

func init() {
	config := sdk.GetConfig()
	config.SetBech32PrefixForValidator("lambvaloper", "lambvaloperpub")
	config.Seal()
}

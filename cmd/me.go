package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/Earth-Online/fofa-api/client"
	"github.com/google/subcommands"
)

type MeCmd struct {
}

func (*MeCmd) Name() string {
	return "me"
}

func (*MeCmd) Synopsis() string {
	return "get you account info"
}

func (*MeCmd) Usage() string {
	return "me"
}

func (c *MeCmd) SetFlags(f *flag.FlagSet) {
}

func (c *MeCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	user := client.NewUser(*email, *key)
	err := user.Me()
	if err != nil {
		fmt.Print(err)
		return subcommands.ExitFailure
	}
	fmt.Printf("%+v\n", user.UserInfo)
	return subcommands.ExitSuccess
}

package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/Earth-Online/fofa-api/client"
	"github.com/google/subcommands"
)

type MsgCmd struct {
}

func (*MsgCmd) Name() string {
	return "msg"
}

func (*MsgCmd) Synopsis() string {
	return "get you notice message"
}

func (*MsgCmd) Usage() string {
	return "msg"
}

func (c *MsgCmd) SetFlags(f *flag.FlagSet) {

}

func (c *MsgCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	user := client.NewUser(*email, *key)
	msg, err := user.GetMessages()
	if err != nil {
		fmt.Print(err)
		return subcommands.ExitFailure
	}
	for _, value := range msg {
		fmt.Printf("%s %s", value.Time, value.Msg)
	}
	return subcommands.ExitSuccess
}

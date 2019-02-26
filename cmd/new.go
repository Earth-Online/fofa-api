package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/google/subcommands"
	"os"
)

type NewCmd struct {
	name string
}

func (*NewCmd) Name() string {
	return "new"
}

func (*NewCmd) Synopsis() string {
	return "new fofa exploit file"
}

func (*NewCmd) Usage() string {
	return "new --name think5_rce"
}

func (c *NewCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.name, "name", "", "filename")
}

func (c *NewCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	var err error
	if c.name == "" {
		return subcommands.ExitUsageError
	}
	if err != nil {
		fmt.Print(err)
		return subcommands.ExitFailure
	}
	output, err := os.OpenFile(c.name+".rb", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Print(err)
		return subcommands.ExitFailure
	}
	_, err = output.WriteString(Template)
	if err != nil {
		fmt.Print(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

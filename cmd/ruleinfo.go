package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/Earth-Online/fofa-api/client"
	"github.com/google/subcommands"
)

type RuleInfoCmd struct {
	name string
}

func (*RuleInfoCmd) Name() string {
	return "ruleinfo"
}

func (*RuleInfoCmd) Synopsis() string {
	return "get rule info"
}

func (*RuleInfoCmd) Usage() string {
	return "ruleinfo --name [rulename]"
}

func (c *RuleInfoCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.name, "name", "", "rulename")
}

func (c *RuleInfoCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if c.name == "" {
		return subcommands.ExitUsageError
	}
	user := client.NewUser(*email, *key)
	info, err := user.GetRuleInfo(c.name)
	if err != nil {
		fmt.Print(err)
		return subcommands.ExitFailure
	}
	fmt.Printf("%s %s", info.Rule, info.Url)
	return subcommands.ExitSuccess
}

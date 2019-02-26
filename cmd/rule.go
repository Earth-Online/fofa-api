package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/Earth-Online/fofa-api/client"
	"github.com/google/subcommands"
)

type RuleCmd struct {
	keyword string
}

func (*RuleCmd) Name() string {
	return "rule"
}

func (*RuleCmd) Synopsis() string {
	return "search rule"
}

func (*RuleCmd) Usage() string {
	return "rule --key [keyword]"
}

func (c *RuleCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.keyword, "key", "", "search keyword")
}

func (c *RuleCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if c.keyword == "" {
		return subcommands.ExitUsageError
	}
	user := client.NewUser(*email, *key)
	rules, err := user.SearchRule(c.keyword)
	if err != nil {
		fmt.Print(err)
		return subcommands.ExitFailure
	}
	if rules.GetErrorMsg() != nil {
		fmt.Print(rules.GetErrorMsg())
		return subcommands.ExitFailure
	}
	fmt.Print("-------search result\n")
	for _, value := range rules.Data {
		fmt.Printf("%s\n", value)
	}
	return subcommands.ExitSuccess
}

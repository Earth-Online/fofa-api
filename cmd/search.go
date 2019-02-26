package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/Earth-Online/fofa-api/client"
	"github.com/google/subcommands"
	"os"
)

type SearchCmd struct {
	keyword    string
	limit      int
	page       int
	outputfile string
}

func (*SearchCmd) Name() string {
	return "search"
}

func (*SearchCmd) Synopsis() string {
	return "search fofa"
}

func (*SearchCmd) Usage() string {
	return "search --keyword [keyword] --limit 100 --output test.txt"
}

func (c *SearchCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&c.keyword, "key", "", "search keyword")
	f.IntVar(&c.limit, "limit", client.MemberLimit, "search result num. member free 100")
	f.IntVar(&c.page, "page", 1, "search page num")
	f.StringVar(&c.outputfile, "output", "output.txt", "output filename")
}

func (c *SearchCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if c.keyword == "" || c.limit <= 0 || c.page <= 0 {
		return subcommands.ExitUsageError
	}
	user := client.NewUser(*email, *key)
	var err error
	result, err := user.Search(c.keyword, "", c.page, c.limit)
	if err != nil {
		fmt.Print(err)
		return subcommands.ExitFailure
	}
	if result.GetErrorMsg() != nil {
		fmt.Print(result.GetErrorMsg())
		return subcommands.ExitFailure
	}
	fmt.Printf("total hava %d query num", result.Size)
	var output *os.File
	if c.outputfile != "" {
		output, err = os.OpenFile(c.outputfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
		if err != nil {
			fmt.Print(err)
			return subcommands.ExitFailure
		}
	} else {
		output = os.Stdout
	}
	for _, value := range result.Results {
		_, err = output.WriteString(fmt.Sprintf("%s", value))
		if err != nil {
			fmt.Print(err)
			return subcommands.ExitFailure
		}
	}
	return subcommands.ExitSuccess
}

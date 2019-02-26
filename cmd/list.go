package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/Earth-Online/fofa-api/client"
	"github.com/google/subcommands"
)

type ListCmd struct {
}

func (*ListCmd) SetFlags(*flag.FlagSet) {

}

func (*ListCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	user := client.NewUser(*email, *key)
	pocs, err := user.GetPoc()
	if err != nil {
		fmt.Print(err)
		return subcommands.ExitFailure
	}
	if pocs.GetErrorMsg() != nil {
		fmt.Print(pocs.GetErrorMsg())
		return subcommands.ExitFailure
	}

	fmt.Printf("------%s\n", client.PUBLISH)
	for _, value := range pocs.Published {
		fmt.Printf("%s-%s\n", value.Name, value.Filename)
	}
	fmt.Printf("------%s\n", client.PEND)
	for _, value := range pocs.Pending {
		fmt.Printf("%s-%s\n", value.Name, value.Filename)
	}
	fmt.Printf("------%s\n", client.PURCHAS)
	for _, value := range pocs.Purchased {
		fmt.Printf("%s-%s\n", value.Name, value.Filename)
	}
	fmt.Printf("------%s\n", client.FAIL)
	for _, value := range pocs.Failed {
		fmt.Printf("%s-%s\n", value.Name, value.Filename)
	}
	return subcommands.ExitSuccess
}

func (*ListCmd) Name() string { return "list" }

func (*ListCmd) Synopsis() string { return "List all your own poc" }

func (*ListCmd) Usage() string {
	return `list:
  .List all your own poc`
}

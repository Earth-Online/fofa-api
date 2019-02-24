package main

import (
	"context"
	"github.com/Earth-Online/fofa-api/cmd"
	"github.com/google/subcommands"
	"os"
)

func main() {
	cfr := subcommands.NewCommander(cmd.FlagSet, "fofa")
	//cfr.Register(subcommands.HelpCommand(), "")
	//cfr.Register(subcommands.FlagsCommand(), "")
	//cfr.Register(subcommands.CommandsCommand(), "")
	cfr.Register(&cmd.ListCmd{}, "")
	cfr.Register(&cmd.DownloadCmd{}, "")
	cfr.Register(&cmd.MeCmd{}, "")
	_ = cmd.FlagSet.Parse(os.Args[1:])
	ctx := context.Background()
	os.Exit(int(cfr.Execute(ctx)))
}

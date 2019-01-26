package cmd

import (
	"flag"
	"os"
)

var (
	email   *string
	key     *string
	FlagSet = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	email = FlagSet.String("email", os.Getenv("FOFA_EMAIL"), "login email")
	key = FlagSet.String("key", os.Getenv("FOFA_KEY"), "login key")
}

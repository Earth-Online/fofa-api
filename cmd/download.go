package cmd

import (
	"context"
	"flag"
	"fmt"
	"github.com/blue-bird1/fofa-api/client"
	"github.com/google/subcommands"
	"os"
)

type DownloadCmd struct {
	filename   string
	outputfile string
}

func (*DownloadCmd) Name() string {
	return "download"
}

func (*DownloadCmd) Synopsis() string {
	return "download you poc file"
}

func (*DownloadCmd) Usage() string {
	return "download -file [filename]"
}

func (d *DownloadCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&d.filename, "filename", "", "download filename")
	f.StringVar(&d.outputfile, "output", "", "output filename")
}

func (d *DownloadCmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if d.filename == "" {
		return subcommands.ExitUsageError
	}
	if d.outputfile == "" {
		d.outputfile = d.filename
	}
	user := client.NewUser(*email, *key)
	code, err := user.GetPocCode(d.filename)
	if err != nil {
		fmt.Print(err)
		return subcommands.ExitFailure
	}
	output, err := os.OpenFile(d.outputfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Print(err)
		return subcommands.ExitFailure
	}
	_, err = output.WriteString(code)
	if err != nil {
		fmt.Print(err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess

}

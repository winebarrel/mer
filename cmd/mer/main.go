package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/shopspring/decimal"
	"github.com/winebarrel/mer"
)

var (
	version string
)

type Options struct {
	From string `arg:"" help:"Exchange source currency code."`
	To   string `arg:"" help:"Exchange destination currency code."`
	Src  string `arg:"" optional:"" help:"Exchange source."`
}

func parseArgs() *Options {
	var CLI struct {
		Options
		Version kong.VersionFlag
	}

	parser := kong.Must(&CLI, kong.Vars{"version": version})
	parser.Model.HelpFlag.Help = "Show help."
	_, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	return &CLI.Options
}

func main() {
	options := parseArgs()

	if options.Src == "" {
		bs, err := io.ReadAll(os.Stdin)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		options.Src = strings.TrimSpace(string(bs))
	}

	src, err := decimal.NewFromString(options.Src)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	dst, err := mer.Exchange(options.From, options.To, src)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(dst)
}

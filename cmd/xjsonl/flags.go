package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var version string

type Flags struct {
	Sep  string
	Keys []string
	File io.ReadCloser
}

func parseArgs() (flags *Flags, err error) {
	flags = &Flags{}
	flag.StringVar(&flags.Sep, "sep", ",", "line separator")
	keys := flag.String("keys", "", "JSON object keys")
	argVersion := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if flag.NArg() == 0 {
		flags.File = os.Stdin
	} else if flag.NArg() == 1 {
		flags.File, err = os.OpenFile(flag.Arg(0), os.O_RDONLY, 0)

		if err != nil {
			return
		}
	} else {
		printUsageAndExit()
	}

	if *argVersion {
		printVersionAndEixt()
	}

	if flags.Sep == "" {
		printErrorAndExit("'-sep' is required")
	}

	flags.Sep = strings.Replace(flags.Sep, "\\t", "\t", -1)

	if *keys != "" {
		flags.Keys = strings.Split(*keys, ",")
	}

	return
}

func printUsageAndExit() {
	fmt.Fprintf(os.Stderr, "Usage: %s\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func printVersionAndEixt() {
	fmt.Fprintln(os.Stderr, version)
	os.Exit(0)
}

func printErrorAndExit(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

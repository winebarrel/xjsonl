package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var version string

type flags struct {
	sep  string
	keys []string
	file io.ReadCloser
}

func parseArgs() (f *flags, err error) {
	f = &flags{}
	flag.StringVar(&f.sep, "sep", ",", "line separator")
	keys := flag.String("keys", "", "JSON object keys")
	argVersion := flag.Bool("version", false, "Print version and exit")
	flag.Parse()

	if flag.NArg() == 0 {
		f.file = os.Stdin
	} else if flag.NArg() == 1 {
		f.file, err = os.OpenFile(flag.Arg(0), os.O_RDONLY, 0)

		if err != nil {
			return
		}
	} else {
		printUsageAndExit()
	}

	if *argVersion {
		printVersionAndEixt()
	}

	if f.sep == "" {
		printErrorAndExit("'-sep' is required")
	}

	f.sep = strings.Replace(f.sep, "\\t", "\t", -1)

	if *keys != "" {
		f.keys = strings.Split(*keys, ",")
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

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
	sep        string
	keys       []string
	withHeader bool
	file       io.ReadCloser
}

func parseArgs() (f *flags, err error) {
	f = &flags{}
	flag.StringVar(&f.sep, "sep", ",", "line separator. not split if empty")
	keys := flag.String("keys", "", "json object keys")
	flag.BoolVar(&f.withHeader, "with-header", false, "consider the first line as a header")
	argVersion := flag.Bool("version", false, "print version and exit")
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

// func printErrorAndExit(msg string) {
// 	fmt.Fprintln(os.Stderr, msg)
// 	os.Exit(1)
// }

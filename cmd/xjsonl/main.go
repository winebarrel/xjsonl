package main

import (
	"fmt"
	"log"

	xjsonl "github.com/winebarrel/xjsonl"
)

func init() {
	log.SetFlags(0)
}

func main() {
	flags, err := parseArgs()

	if err != nil {
		log.Fatal(err)
	}

	defer flags.file.Close()

	err = xjsonl.EachJsonLine(flags.file, flags.sep, flags.keys, func(line string) {
		fmt.Println(line)
	})

	if err != nil {
		log.Fatal(err)
	}
}

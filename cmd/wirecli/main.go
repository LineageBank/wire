// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

/*
 * to run from cli, make sure terminal is in cmd/wirecli
 * run: go run main.go reformat.go describe.go -reformat=json the_fedwire_file_name.txt
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/moov-io/wire"
)

var (
	flagReformat = flag.String("reformat", "", "Reformat an incoming Wire file to json format")
	programName  = filepath.Base(os.Args[0])
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage (%s):\n", wire.Version)
		fmt.Printf("   usage: %s [<flags>] <files>", programName)
		fmt.Println("")
		fmt.Println("")
		fmt.Println("Commands: ")
		fmt.Printf("  %s -reformat=json wire.txt", programName)
		fmt.Println("  [ Convert an incoming Wire file into json format ]")

		fmt.Println("")
		fmt.Println("Flags: ")
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	args := flag.Args()

	// pick our command to do
	switch {

	case *flagReformat != "" && len(args) == 1:
		if err := reformat(*flagReformat, args[0]); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}

	default:
		if err := dumpFiles(args); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}
	}
}

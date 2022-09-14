// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"

	"github.com/moov-io/wire"
)

func dumpFiles(paths []string) error {
	var files []*wire.File
	for i := range paths {
		f, err := readFile(paths[i])
		if err != nil {
			fmt.Printf("WARN: problem reading %s:\n %v\n\n", paths[i], err)
		}
		files = append(files, f)
	}

	for i := range files {
		if i > 0 && len(files) > 1 {
			fmt.Println("") // extra newline between multiple ACH files
		} else {
			fmt.Printf("nil Wire file in position %d\n", i)
		}
	}

	return nil
}

func readFile(path string) (*wire.File, error) {
	fd, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s: %v", path, err)
	}
	defer fd.Close()

	f, err := wire.NewReader(fd).Read()
	return &f, err
}

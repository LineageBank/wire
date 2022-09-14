// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/moov-io/wire"
)

func reformat(as string, filepath string) error {
	if _, err := os.Stat(filepath); err != nil {
		return err
	}

	file, err := readIncomingFile(filepath)
	if err != nil {
		return err
	}

	switch as {
	// case "json":
	// 	w := wire.NewWriter(os.Stdout)
	// 	if err := w.Write(file); err != nil {
	// 		return err
	// 	}

	case "json":
		// if err := json.RawMessage(os.Stdout).Encode(file); err != nil {
		if err := json.NewEncoder(os.Stdout).Encode(file); err != nil {
			return err
		}

	default:
		return fmt.Errorf("unknown format %s", as)
	}
	return nil
}

func readIncomingFile(path string) (*wire.File, error) {
	file, err := readJsonFile(path)
	if file != nil && err == nil {
		return file, nil
	}

	return nil, fmt.Errorf("unable to read %s:\n %v", path, err)
}

func readJsonFile(path string) (*wire.File, error) {

	fd, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("problem opening %s: %v", path, err)
	}
	defer fd.Close()

	// fd, err := os.ReadFile(path)
	// if err != nil {
	// 	return nil, fmt.Errorf("problem reading %s: %v", path, err)
	// }
	// str := string(fd)
	// str2 := json.Encoder(str)
	// println(str2)

	bs, err := io.ReadAll(fd)
	if err != nil {
		return nil, fmt.Errorf("problem reading %s: %v", path, err)
	}

	return wire.FileFromJSON(bs)
}

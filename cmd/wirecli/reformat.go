// Copyright 2019 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/moov-io/wire"
)

func reformat(as string, filepath string) error {
	if _, err := os.Stat(filepath); err != nil {
		return err
	}

	file, err := readWireFile(filepath)
	if err != nil {
		return err
	}

	switch as {

	case "json":
		if err := json.NewEncoder(os.Stdout).Encode(file); err != nil {
			return err
		}

	default:
		return fmt.Errorf("unknown format %s", as)
	}
	return nil
}

func readWireFile(path string) (*wire.File, error) {

	fd, err := os.Open(path)
	if err != nil {
		fmt.Printf("problem opening %s: %v", path, err)
	}
	defer fd.Close()

	r4 := bufio.NewReader(fd)
	b4, err := r4.Peek(r4.Size())
	if err != nil {
		if err == io.EOF {
			fmt.Println("")
		} else {
			fmt.Printf("problem opening %s: %v", path, err)
		}
	}

	parsed, err := parseContents(string(b4))
	if err != nil {
		fmt.Printf("Unable to parse file %s: %v", path, err)
	}

	pretty, err := prettyJson(parsed)
	if err != nil {
		fmt.Printf("unable to convert wire file to json %s\n", err)
	}

	return wire.FileFromJSON(pretty)

}

func parseContents(input string) (string, error) {
	r := strings.NewReader(input)
	file, err := wire.NewReader(r).Read()
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(file); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func prettyJson(input string) ([]byte, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return nil, err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return nil, err
	}
	return pretty, nil
}

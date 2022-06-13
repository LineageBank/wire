// Copyright 2020 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package wire

import (
	"encoding/json"
	"strings"
	"unicode/utf8"
)

// FIBeneficiaryFI is the financial institution beneficiary financial institution
type FIBeneficiaryFI struct {
	// tag
	tag string
	// Financial Institution
	FIToFI FIToFI `json:"fiToFI,omitempty"`

	// validator is composed for data validation
	validator
	// converters is composed for WIRE to GoLang Converters
	converters
}

// NewFIBeneficiaryFI returns a new FIBeneficiaryFI
func NewFIBeneficiaryFI() *FIBeneficiaryFI {
	fibfi := &FIBeneficiaryFI{
		tag: TagFIBeneficiaryFI,
	}
	return fibfi
}

// Parse takes the input string and parses the FIBeneficiaryFI values
//
// Parse provides no guarantee about all fields being filled in. Callers should make a Validate() call to confirm
// successful parsing and data validity.
func (fibfi *FIBeneficiaryFI) Parse(record string) error {
	if utf8.RuneCountInString(record) < 6 {
		return NewTagMinLengthErr(6, len(record))
	}

	fibfi.tag = record[:6]
	length := 6

	value, read, err := fibfi.parseVariableStringField(record[length:], 30)
	if err != nil {
		return fieldError("LineOne", err)
	}
	fibfi.FIToFI.LineOne = value
	length += read

	value, read, err = fibfi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineTwo", err)
	}
	fibfi.FIToFI.LineTwo = value
	length += read

	value, read, err = fibfi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineThree", err)
	}
	fibfi.FIToFI.LineThree = value
	length += read

	value, read, err = fibfi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFour", err)
	}
	fibfi.FIToFI.LineFour = value
	length += read

	value, read, err = fibfi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineFive", err)
	}
	fibfi.FIToFI.LineFive = value
	length += read

	value, read, err = fibfi.parseVariableStringField(record[length:], 33)
	if err != nil {
		return fieldError("LineSix", err)
	}
	fibfi.FIToFI.LineSix = value
	length += read

	if !fibfi.verifyDataWithReadLength(record, length) {
		return NewTagMaxLengthErr()
	}

	return nil
}

func (fibfi *FIBeneficiaryFI) UnmarshalJSON(data []byte) error {
	type Alias FIBeneficiaryFI
	aux := struct {
		*Alias
	}{
		(*Alias)(fibfi),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	fibfi.tag = TagFIBeneficiaryFI
	return nil
}

// String writes FIBeneficiaryFI
func (fibfi *FIBeneficiaryFI) String(options ...bool) string {
	var buf strings.Builder
	buf.Grow(201)

	buf.WriteString(fibfi.tag)
	buf.WriteString(fibfi.LineOneField(options...))
	buf.WriteString(fibfi.LineTwoField(options...))
	buf.WriteString(fibfi.LineThreeField(options...))
	buf.WriteString(fibfi.LineFourField(options...))
	buf.WriteString(fibfi.LineFiveField(options...))
	buf.WriteString(fibfi.LineSixField(options...))

	if fibfi.parseFirstOption(options) {
		return fibfi.stripDelimiters(buf.String())
	} else {
		return buf.String()
	}
}

// Validate performs WIRE format rule checks on FIBeneficiaryFI and returns an error if not Validated
// The first error encountered is returned and stops that parsing.
func (fibfi *FIBeneficiaryFI) Validate() error {
	if fibfi.tag != TagFIBeneficiaryFI {
		return fieldError("tag", ErrValidTagForType, fibfi.tag)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineOne); err != nil {
		return fieldError("LineOne", err, fibfi.FIToFI.LineOne)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineTwo); err != nil {
		return fieldError("LineTwo", err, fibfi.FIToFI.LineTwo)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineThree); err != nil {
		return fieldError("LineThree", err, fibfi.FIToFI.LineThree)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineFour); err != nil {
		return fieldError("LineFour", err, fibfi.FIToFI.LineFour)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineFive); err != nil {
		return fieldError("LineFive", err, fibfi.FIToFI.LineFive)
	}
	if err := fibfi.isAlphanumeric(fibfi.FIToFI.LineSix); err != nil {
		return fieldError("LineSix", err, fibfi.FIToFI.LineSix)
	}
	return nil
}

// LineOneField gets a string of the LineOne field
func (fibfi *FIBeneficiaryFI) LineOneField(options ...bool) string {
	return fibfi.alphaVariableField(fibfi.FIToFI.LineOne, 30, fibfi.parseFirstOption(options))
}

// LineTwoField gets a string of the LineTwo field
func (fibfi *FIBeneficiaryFI) LineTwoField(options ...bool) string {
	return fibfi.alphaVariableField(fibfi.FIToFI.LineTwo, 33, fibfi.parseFirstOption(options))
}

// LineThreeField gets a string of the LineThree field
func (fibfi *FIBeneficiaryFI) LineThreeField(options ...bool) string {
	return fibfi.alphaVariableField(fibfi.FIToFI.LineThree, 33, fibfi.parseFirstOption(options))
}

// LineFourField gets a string of the LineFour field
func (fibfi *FIBeneficiaryFI) LineFourField(options ...bool) string {
	return fibfi.alphaVariableField(fibfi.FIToFI.LineFour, 33, fibfi.parseFirstOption(options))
}

// LineFiveField gets a string of the LineFive field
func (fibfi *FIBeneficiaryFI) LineFiveField(options ...bool) string {
	return fibfi.alphaVariableField(fibfi.FIToFI.LineFive, 33, fibfi.parseFirstOption(options))
}

// LineSixField gets a string of the LineSix field
func (fibfi *FIBeneficiaryFI) LineSixField(options ...bool) string {
	return fibfi.alphaVariableField(fibfi.FIToFI.LineSix, 33, fibfi.parseFirstOption(options))
}

// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type FitBaseUnit uint16

const (
	FitBaseUnitOther    FitBaseUnit = 0
	FitBaseUnitKilogram FitBaseUnit = 1
	FitBaseUnitPound    FitBaseUnit = 2
	FitBaseUnitInvalid  FitBaseUnit = 0xFFFF
)

func (f FitBaseUnit) Uint16() uint16 { return uint16(f) }

func (f FitBaseUnit) String() string {
	switch f {
	case FitBaseUnitOther:
		return "other"
	case FitBaseUnitKilogram:
		return "kilogram"
	case FitBaseUnitPound:
		return "pound"
	default:
		return "FitBaseUnitInvalid(" + strconv.FormatUint(uint64(f), 10) + ")"
	}
}

// FromString parse string into FitBaseUnit constant it's represent, return FitBaseUnitInvalid if not found.
func FitBaseUnitFromString(s string) FitBaseUnit {
	switch s {
	case "other":
		return FitBaseUnitOther
	case "kilogram":
		return FitBaseUnitKilogram
	case "pound":
		return FitBaseUnitPound
	default:
		return FitBaseUnitInvalid
	}
}

// List returns all constants.
func ListFitBaseUnit() []FitBaseUnit {
	return []FitBaseUnit{
		FitBaseUnitOther,
		FitBaseUnitKilogram,
		FitBaseUnitPound,
	}
}
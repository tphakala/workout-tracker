// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type MesgCount byte

const (
	MesgCountNumPerFile     MesgCount = 0
	MesgCountMaxPerFile     MesgCount = 1
	MesgCountMaxPerFileType MesgCount = 2
	MesgCountInvalid        MesgCount = 0xFF
)

func (m MesgCount) Byte() byte { return byte(m) }

func (m MesgCount) String() string {
	switch m {
	case MesgCountNumPerFile:
		return "num_per_file"
	case MesgCountMaxPerFile:
		return "max_per_file"
	case MesgCountMaxPerFileType:
		return "max_per_file_type"
	default:
		return "MesgCountInvalid(" + strconv.Itoa(int(m)) + ")"
	}
}

// FromString parse string into MesgCount constant it's represent, return MesgCountInvalid if not found.
func MesgCountFromString(s string) MesgCount {
	switch s {
	case "num_per_file":
		return MesgCountNumPerFile
	case "max_per_file":
		return MesgCountMaxPerFile
	case "max_per_file_type":
		return MesgCountMaxPerFileType
	default:
		return MesgCountInvalid
	}
}

// List returns all constants.
func ListMesgCount() []MesgCount {
	return []MesgCount{
		MesgCountNumPerFile,
		MesgCountMaxPerFile,
		MesgCountMaxPerFileType,
	}
}
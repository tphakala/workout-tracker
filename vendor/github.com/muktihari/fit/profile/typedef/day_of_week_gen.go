// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typedef

import (
	"strconv"
)

type DayOfWeek byte

const (
	DayOfWeekSunday    DayOfWeek = 0
	DayOfWeekMonday    DayOfWeek = 1
	DayOfWeekTuesday   DayOfWeek = 2
	DayOfWeekWednesday DayOfWeek = 3
	DayOfWeekThursday  DayOfWeek = 4
	DayOfWeekFriday    DayOfWeek = 5
	DayOfWeekSaturday  DayOfWeek = 6
	DayOfWeekInvalid   DayOfWeek = 0xFF
)

func (d DayOfWeek) Byte() byte { return byte(d) }

func (d DayOfWeek) String() string {
	switch d {
	case DayOfWeekSunday:
		return "sunday"
	case DayOfWeekMonday:
		return "monday"
	case DayOfWeekTuesday:
		return "tuesday"
	case DayOfWeekWednesday:
		return "wednesday"
	case DayOfWeekThursday:
		return "thursday"
	case DayOfWeekFriday:
		return "friday"
	case DayOfWeekSaturday:
		return "saturday"
	default:
		return "DayOfWeekInvalid(" + strconv.Itoa(int(d)) + ")"
	}
}

// FromString parse string into DayOfWeek constant it's represent, return DayOfWeekInvalid if not found.
func DayOfWeekFromString(s string) DayOfWeek {
	switch s {
	case "sunday":
		return DayOfWeekSunday
	case "monday":
		return DayOfWeekMonday
	case "tuesday":
		return DayOfWeekTuesday
	case "wednesday":
		return DayOfWeekWednesday
	case "thursday":
		return DayOfWeekThursday
	case "friday":
		return DayOfWeekFriday
	case "saturday":
		return DayOfWeekSaturday
	default:
		return DayOfWeekInvalid
	}
}

// List returns all constants.
func ListDayOfWeek() []DayOfWeek {
	return []DayOfWeek{
		DayOfWeekSunday,
		DayOfWeekMonday,
		DayOfWeekTuesday,
		DayOfWeekWednesday,
		DayOfWeekThursday,
		DayOfWeekFriday,
		DayOfWeekSaturday,
	}
}
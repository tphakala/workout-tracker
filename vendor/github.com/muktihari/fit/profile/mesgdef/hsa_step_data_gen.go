// Code generated by internal/cmd/fitgen/main.go. DO NOT EDIT.

// Copyright 2023 The FIT SDK for Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mesgdef

import (
	"github.com/muktihari/fit/factory"
	"github.com/muktihari/fit/internal/sliceutil"
	"github.com/muktihari/fit/kit/datetime"
	"github.com/muktihari/fit/profile/basetype"
	"github.com/muktihari/fit/profile/typedef"
	"github.com/muktihari/fit/proto"
	"time"
)

// HsaStepData is a HsaStepData message.
//
// Note: The order of the fields is optimized using a memory alignment algorithm.
// Do not rely on field indices, such as when using reflection.
type HsaStepData struct {
	Timestamp          time.Time // Units: s
	Steps              []uint32  // Array: [N]; Units: steps; Total step sum
	ProcessingInterval uint16    // Units: s; Processing interval length in seconds. File start: 0xFFFFFFEF File stop: 0xFFFFFFEE

	UnknownFields   []proto.Field          // UnknownFields are fields that are exist but they are not defined in Profile.xlsx
	DeveloperFields []proto.DeveloperField // DeveloperFields are custom data fields [Added since protocol version 2.0]
}

// NewHsaStepData creates new HsaStepData struct based on given mesg.
// If mesg is nil, it will return HsaStepData with all fields being set to its corresponding invalid value.
func NewHsaStepData(mesg *proto.Message) *HsaStepData {
	vals := [254]proto.Value{}

	var unknownFields []proto.Field
	var developerFields []proto.DeveloperField
	if mesg != nil {
		arr := pool.Get().(*[poolsize]proto.Field)
		unknownFields = arr[:0]
		for i := range mesg.Fields {
			if mesg.Fields[i].Num > 253 || mesg.Fields[i].Name == factory.NameUnknown {
				unknownFields = append(unknownFields, mesg.Fields[i])
				continue
			}
			vals[mesg.Fields[i].Num] = mesg.Fields[i].Value
		}
		unknownFields = sliceutil.Clone(unknownFields)
		*arr = [poolsize]proto.Field{}
		pool.Put(arr)
		developerFields = mesg.DeveloperFields
	}

	return &HsaStepData{
		Timestamp:          datetime.ToTime(vals[253].Uint32()),
		ProcessingInterval: vals[0].Uint16(),
		Steps:              vals[1].SliceUint32(),

		UnknownFields:   unknownFields,
		DeveloperFields: developerFields,
	}
}

// ToMesg converts HsaStepData into proto.Message. If options is nil, default options will be used.
func (m *HsaStepData) ToMesg(options *Options) proto.Message {
	if options == nil {
		options = defaultOptions
	} else if options.Factory == nil {
		options.Factory = factory.StandardFactory()
	}

	fac := options.Factory

	arr := pool.Get().(*[poolsize]proto.Field)
	fields := arr[:0]

	mesg := proto.Message{Num: typedef.MesgNumHsaStepData}

	if !m.Timestamp.Before(datetime.Epoch()) {
		field := fac.CreateField(mesg.Num, 253)
		field.Value = proto.Uint32(uint32(m.Timestamp.Sub(datetime.Epoch()).Seconds()))
		fields = append(fields, field)
	}
	if m.ProcessingInterval != basetype.Uint16Invalid {
		field := fac.CreateField(mesg.Num, 0)
		field.Value = proto.Uint16(m.ProcessingInterval)
		fields = append(fields, field)
	}
	if m.Steps != nil {
		field := fac.CreateField(mesg.Num, 1)
		field.Value = proto.SliceUint32(m.Steps)
		fields = append(fields, field)
	}

	for i := range m.UnknownFields {
		fields = append(fields, m.UnknownFields[i])
	}

	mesg.Fields = make([]proto.Field, len(fields))
	copy(mesg.Fields, fields)
	*arr = [poolsize]proto.Field{}
	pool.Put(arr)

	mesg.DeveloperFields = m.DeveloperFields

	return mesg
}

// TimestampUint32 returns Timestamp in uint32 (seconds since FIT's epoch) instead of time.Time.
func (m *HsaStepData) TimestampUint32() uint32 { return datetime.ToUint32(m.Timestamp) }

// SetTimestamp sets Timestamp value.
//
// Units: s
func (m *HsaStepData) SetTimestamp(v time.Time) *HsaStepData {
	m.Timestamp = v
	return m
}

// SetProcessingInterval sets ProcessingInterval value.
//
// Units: s; Processing interval length in seconds. File start: 0xFFFFFFEF File stop: 0xFFFFFFEE
func (m *HsaStepData) SetProcessingInterval(v uint16) *HsaStepData {
	m.ProcessingInterval = v
	return m
}

// SetSteps sets Steps value.
//
// Array: [N]; Units: steps; Total step sum
func (m *HsaStepData) SetSteps(v []uint32) *HsaStepData {
	m.Steps = v
	return m
}

// SetUnknownFields sets UnknownFields (fields that are exist but they are not defined in Profile.xlsx)
func (m *HsaStepData) SetUnknownFields(unknownFields ...proto.Field) *HsaStepData {
	m.UnknownFields = unknownFields
	return m
}

// SetDeveloperFields sets DeveloperFields.
func (m *HsaStepData) SetDeveloperFields(developerFields ...proto.DeveloperField) *HsaStepData {
	m.DeveloperFields = developerFields
	return m
}
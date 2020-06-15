// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/enums/v1/common.proto

package enums

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type EncodingType int32

const (
	ENCODING_TYPE_UNSPECIFIED EncodingType = 0
	ENCODING_TYPE_PROTO3      EncodingType = 1
	ENCODING_TYPE_JSON        EncodingType = 2
)

var EncodingType_name = map[int32]string{
	0: "Unspecified",
	1: "Proto3",
	2: "Json",
}

var EncodingType_value = map[string]int32{
	"Unspecified": 0,
	"Proto3":      1,
	"Json":        2,
}

func (EncodingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_03a8b132cb70d36d, []int{0}
}

type IndexedValueType int32

const (
	INDEXED_VALUE_TYPE_UNSPECIFIED IndexedValueType = 0
	INDEXED_VALUE_TYPE_STRING      IndexedValueType = 1
	INDEXED_VALUE_TYPE_KEYWORD     IndexedValueType = 2
	INDEXED_VALUE_TYPE_INT         IndexedValueType = 3
	INDEXED_VALUE_TYPE_DOUBLE      IndexedValueType = 4
	INDEXED_VALUE_TYPE_BOOL        IndexedValueType = 5
	INDEXED_VALUE_TYPE_DATETIME    IndexedValueType = 6
)

var IndexedValueType_name = map[int32]string{
	0: "Unspecified",
	1: "String",
	2: "Keyword",
	3: "Int",
	4: "Double",
	5: "Bool",
	6: "Datetime",
}

var IndexedValueType_value = map[string]int32{
	"Unspecified": 0,
	"String":      1,
	"Keyword":     2,
	"Int":         3,
	"Double":      4,
	"Bool":        5,
	"Datetime":    6,
}

func (IndexedValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_03a8b132cb70d36d, []int{1}
}

func init() {
	proto.RegisterEnum("temporal.enums.v1.EncodingType", EncodingType_name, EncodingType_value)
	proto.RegisterEnum("temporal.enums.v1.IndexedValueType", IndexedValueType_name, IndexedValueType_value)
}

func init() { proto.RegisterFile("temporal/enums/v1/common.proto", fileDescriptor_03a8b132cb70d36d) }

var fileDescriptor_03a8b132cb70d36d = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x31, 0x4f, 0xf2, 0x40,
	0x1c, 0xc6, 0x7b, 0xbc, 0xef, 0xcb, 0x70, 0xaf, 0xc3, 0x79, 0x21, 0x28, 0x10, 0xfe, 0x26, 0x8e,
	0x44, 0xdb, 0x10, 0x46, 0x27, 0x4a, 0x4f, 0x72, 0x8a, 0xbd, 0x06, 0x0a, 0x8a, 0x4b, 0x83, 0xd0,
	0x10, 0x12, 0xda, 0x23, 0x08, 0x44, 0x37, 0x37, 0x57, 0x3f, 0x86, 0x1f, 0xc5, 0x91, 0x91, 0x51,
	0x8e, 0xc5, 0xc9, 0xf0, 0x11, 0x8c, 0x4d, 0x30, 0x12, 0xeb, 0xf6, 0xe4, 0x7e, 0xcf, 0xfd, 0x9f,
	0xe1, 0x87, 0x61, 0xe2, 0x07, 0x23, 0x39, 0xee, 0x0c, 0x0d, 0x3f, 0x9c, 0x06, 0xb7, 0xc6, 0xac,
	0x68, 0x74, 0x65, 0x10, 0xc8, 0x50, 0x1f, 0x8d, 0xe5, 0x44, 0xd2, 0xdd, 0x0d, 0xd7, 0x23, 0xae,
	0xcf, 0x8a, 0x05, 0x0f, 0xef, 0xb0, 0xb0, 0x2b, 0x7b, 0x83, 0xb0, 0xef, 0xde, 0x8f, 0x7c, 0x9a,
	0xc7, 0x19, 0x66, 0x57, 0x84, 0xc5, 0xed, 0xaa, 0xe7, 0xb6, 0x1d, 0xe6, 0x35, 0xed, 0x86, 0xc3,
	0x2a, 0xfc, 0x94, 0x33, 0x8b, 0x68, 0x74, 0x1f, 0xa7, 0xb6, 0xb1, 0x53, 0x17, 0xae, 0x28, 0x11,
	0x44, 0xd3, 0x98, 0x6e, 0x93, 0xb3, 0x86, 0xb0, 0x49, 0xa2, 0xf0, 0x8e, 0x30, 0xe1, 0x61, 0xcf,
	0xbf, 0xf3, 0x7b, 0xad, 0xce, 0x70, 0xea, 0x47, 0x2b, 0x87, 0x18, 0xb8, 0x6d, 0xb1, 0x2b, 0x66,
	0x79, 0xad, 0x72, 0xad, 0xc9, 0xe2, 0xa6, 0xf2, 0x38, 0x13, 0xd3, 0x69, 0xb8, 0x75, 0x6e, 0x57,
	0x09, 0xa2, 0x80, 0xb3, 0x31, 0xf8, 0x9c, 0xb5, 0x2f, 0x45, 0xdd, 0x22, 0x09, 0x9a, 0xc5, 0xe9,
	0x18, 0xce, 0x6d, 0x97, 0xfc, 0xf9, 0xe5, 0xb4, 0x25, 0x9a, 0x66, 0x8d, 0x91, 0xbf, 0x34, 0x87,
	0xf7, 0x62, 0xb0, 0x29, 0x44, 0x8d, 0xfc, 0xa3, 0x07, 0x38, 0x17, 0xf7, 0xb7, 0xec, 0x32, 0x97,
	0x5f, 0x30, 0x92, 0x34, 0x1f, 0xd1, 0x7c, 0x09, 0xda, 0x62, 0x09, 0xda, 0x7a, 0x09, 0xe8, 0x41,
	0x01, 0x7a, 0x56, 0x80, 0x5e, 0x14, 0xa0, 0xb9, 0x02, 0xf4, 0xaa, 0x00, 0xbd, 0x29, 0xd0, 0xd6,
	0x0a, 0xd0, 0xd3, 0x0a, 0xb4, 0xf9, 0x0a, 0xb4, 0xc5, 0x0a, 0x34, 0x9c, 0x1a, 0x48, 0xfd, 0x87,
	0x1e, 0xf3, 0x7f, 0x25, 0xf2, 0xe7, 0x7c, 0xea, 0x73, 0xd0, 0xf5, 0x51, 0xff, 0x5b, 0x69, 0x20,
	0x8d, 0x4d, 0x3e, 0x8e, 0xfc, 0x7e, 0x59, 0x3f, 0x89, 0xc2, 0x4d, 0x32, 0x7a, 0x2d, 0x7d, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x47, 0x0a, 0x05, 0xed, 0x17, 0x02, 0x00, 0x00,
}

func (x EncodingType) String() string {
	s, ok := EncodingType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x IndexedValueType) String() string {
	s, ok := IndexedValueType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

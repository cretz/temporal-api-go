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
// source: temporal/enums/v1/query.proto

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

type QueryResultType int32

const (
	QUERY_RESULT_TYPE_UNSPECIFIED QueryResultType = 0
	QUERY_RESULT_TYPE_ANSWERED    QueryResultType = 1
	QUERY_RESULT_TYPE_FAILED      QueryResultType = 2
)

var QueryResultType_name = map[int32]string{
	0: "Unspecified",
	1: "Answered",
	2: "Failed",
}

var QueryResultType_value = map[string]int32{
	"Unspecified": 0,
	"Answered":    1,
	"Failed":      2,
}

func (QueryResultType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e9ca087af688e80, []int{0}
}

type QueryRejectCondition int32

const (
	// Unspecified indicates that query should not be rejected.
	QUERY_REJECT_CONDITION_UNSPECIFIED QueryRejectCondition = 0
	// NotOpen indicates that query should be rejected if workflow is not open.
	QUERY_REJECT_CONDITION_NOT_OPEN QueryRejectCondition = 1
	// NotCompletedCleanly indicates that query should be rejected if workflow did not complete cleanly.
	QUERY_REJECT_CONDITION_NOT_COMPLETED_CLEANLY QueryRejectCondition = 2
)

var QueryRejectCondition_name = map[int32]string{
	0: "Unspecified",
	1: "NotOpen",
	2: "NotCompletedCleanly",
}

var QueryRejectCondition_value = map[string]int32{
	"Unspecified":         0,
	"NotOpen":             1,
	"NotCompletedCleanly": 2,
}

func (QueryRejectCondition) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e9ca087af688e80, []int{1}
}

type QueryConsistencyLevel int32

const (
	QUERY_CONSISTENCY_LEVEL_UNSPECIFIED QueryConsistencyLevel = 0
	// Eventual indicates that query should be eventually consistent.
	QUERY_CONSISTENCY_LEVEL_EVENTUAL QueryConsistencyLevel = 1
	// Strong indicates that any events that came before query should be reflected in workflow state before running query.
	QUERY_CONSISTENCY_LEVEL_STRONG QueryConsistencyLevel = 2
)

var QueryConsistencyLevel_name = map[int32]string{
	0: "Unspecified",
	1: "Eventual",
	2: "Strong",
}

var QueryConsistencyLevel_value = map[string]int32{
	"Unspecified": 0,
	"Eventual":    1,
	"Strong":      2,
}

func (QueryConsistencyLevel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0e9ca087af688e80, []int{2}
}

func init() {
	proto.RegisterEnum("temporal.enums.v1.QueryResultType", QueryResultType_name, QueryResultType_value)
	proto.RegisterEnum("temporal.enums.v1.QueryRejectCondition", QueryRejectCondition_name, QueryRejectCondition_value)
	proto.RegisterEnum("temporal.enums.v1.QueryConsistencyLevel", QueryConsistencyLevel_name, QueryConsistencyLevel_value)
}

func init() { proto.RegisterFile("temporal/enums/v1/query.proto", fileDescriptor_0e9ca087af688e80) }

var fileDescriptor_0e9ca087af688e80 = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xb1, 0xae, 0xd3, 0x30,
	0x18, 0x85, 0xe3, 0x0e, 0x0c, 0x5e, 0x08, 0xd6, 0x45, 0x42, 0x88, 0x6b, 0xe0, 0x5e, 0x04, 0x52,
	0x55, 0x12, 0x2a, 0x46, 0xa6, 0xd4, 0xf9, 0x8b, 0x82, 0x8c, 0x93, 0x26, 0x4e, 0x51, 0x58, 0x22,
	0x28, 0x16, 0x0a, 0x6a, 0xe3, 0x92, 0xa4, 0x95, 0xba, 0x31, 0x31, 0x30, 0xb1, 0xf0, 0x0e, 0x3c,
	0x0a, 0x63, 0xc7, 0x8e, 0x34, 0x5d, 0x18, 0xfb, 0x08, 0xa8, 0x81, 0x22, 0xd4, 0xaa, 0x77, 0xfb,
	0xe5, 0xf3, 0xd9, 0xe7, 0x58, 0xff, 0xc1, 0xe7, 0x95, 0x9a, 0x4c, 0x75, 0xf1, 0x66, 0x6c, 0xab,
	0x7c, 0x36, 0x29, 0xed, 0x79, 0xd7, 0xfe, 0x38, 0x53, 0xc5, 0xc2, 0x9a, 0x16, 0xba, 0xd2, 0xe4,
	0xc6, 0x5e, 0xb6, 0x1a, 0xd9, 0x9a, 0x77, 0xdb, 0x05, 0xbe, 0x3e, 0xd8, 0x11, 0xa1, 0x2a, 0x67,
	0xe3, 0x4a, 0x2e, 0xa6, 0x8a, 0xdc, 0xc7, 0xe7, 0x83, 0x18, 0xc2, 0x24, 0x0d, 0x21, 0x8a, 0xb9,
	0x4c, 0x65, 0x12, 0x40, 0x1a, 0x8b, 0x28, 0x00, 0xe6, 0xf5, 0x3d, 0x70, 0x4d, 0x83, 0x50, 0x7c,
	0xfb, 0x18, 0x71, 0x44, 0xf4, 0x0a, 0x42, 0x70, 0x4d, 0x44, 0xee, 0xe0, 0x5b, 0xc7, 0x7a, 0xdf,
	0xf1, 0x38, 0xb8, 0x66, 0xab, 0xfd, 0x0d, 0xe1, 0xb3, 0xbf, 0xa6, 0x1f, 0xd4, 0xa8, 0x62, 0x3a,
	0x7f, 0x97, 0x55, 0x99, 0xce, 0xc9, 0x43, 0x7c, 0xb1, 0xbf, 0xf6, 0x02, 0x98, 0x4c, 0x99, 0x2f,
	0x5c, 0x4f, 0x7a, 0xbe, 0x38, 0xb0, 0xbf, 0xc4, 0x77, 0x4f, 0x70, 0xc2, 0x97, 0xa9, 0x1f, 0x80,
	0x30, 0x11, 0x79, 0x82, 0x3b, 0x57, 0x40, 0xcc, 0x7f, 0x19, 0x70, 0x90, 0xe0, 0xa6, 0x8c, 0x83,
	0x23, 0x78, 0x62, 0xb6, 0xda, 0x5f, 0x10, 0xbe, 0xd9, 0xe4, 0x62, 0x3a, 0x2f, 0xb3, 0xb2, 0x52,
	0xf9, 0x68, 0xc1, 0xd5, 0x5c, 0x8d, 0xc9, 0x23, 0x7c, 0xf9, 0xe7, 0x2d, 0xe6, 0x8b, 0xc8, 0x8b,
	0x24, 0x08, 0x96, 0xa4, 0x1c, 0x86, 0xc0, 0x0f, 0x92, 0x3d, 0xc0, 0xf7, 0x4e, 0x81, 0x30, 0x04,
	0x21, 0x63, 0x87, 0x9b, 0x88, 0x5c, 0x60, 0x7a, 0x8a, 0x8a, 0x64, 0xe8, 0x8b, 0xe7, 0x66, 0xab,
	0xf7, 0x19, 0x2d, 0xd7, 0xd4, 0x58, 0xad, 0xa9, 0xb1, 0x5d, 0x53, 0xf4, 0xa9, 0xa6, 0xe8, 0x7b,
	0x4d, 0xd1, 0x8f, 0x9a, 0xa2, 0x65, 0x4d, 0xd1, 0xcf, 0x9a, 0xa2, 0x5f, 0x35, 0x35, 0xb6, 0x35,
	0x45, 0x5f, 0x37, 0xd4, 0x58, 0x6e, 0xa8, 0xb1, 0xda, 0x50, 0x03, 0x9f, 0x65, 0xda, 0x3a, 0xda,
	0x72, 0x0f, 0x37, 0xdf, 0x0a, 0x76, 0x25, 0x08, 0xd0, 0xeb, 0xce, 0xfb, 0xff, 0x98, 0x4c, 0xdb,
	0xfb, 0xf9, 0x71, 0xd3, 0x92, 0x7f, 0xd5, 0x79, 0xd6, 0x0c, 0x6f, 0xaf, 0x35, 0xa7, 0x4f, 0x7f,
	0x07, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x16, 0x2c, 0xcb, 0x5c, 0x02, 0x00, 0x00,
}

func (x QueryResultType) String() string {
	s, ok := QueryResultType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x QueryRejectCondition) String() string {
	s, ok := QueryRejectCondition_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x QueryConsistencyLevel) String() string {
	s, ok := QueryConsistencyLevel_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
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
// source: temporal/api/sdk/v1/task_complete_metadata.proto

package sdk

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
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

type WorkflowTaskCompletedMetadata struct {
	// Internal flags used by the core SDK. SDKs using flags must comply with the following behavior:
	//
	// During replay:
	//   - If a flag is not recognized (value is too high or not defined), it must fail the workflow
	//     task.
	//   - If a flag is recognized, it is stored in a set of used flags for the run. Code checks for
	//     that flag during and after this WFT are allowed to assume that the flag is present.
	//   - If a code check for a flag does not find the flag in the set of used flags, it must take
	//     the branch corresponding to the absence of that flag.
	//
	// During non-replay execution of new WFTs:
	//   - The SDK is free to use all flags it knows about. It must record any newly-used (IE: not
	//     previously recorded) flags when completing the WFT.
	//
	// SDKs which are too old to even know about this field at all are considered to produce
	// undefined behavior if they replay workflows which used this mechanism.
	//
	// (-- api-linter: core::0141::forbidden-types=disabled
	//
	//	aip.dev/not-precedent: These really shouldn't have negative values. --)
	CoreUsedFlags []uint32 `protobuf:"varint,1,rep,packed,name=core_used_flags,json=coreUsedFlags,proto3" json:"core_used_flags,omitempty"`
	// Flags used by the SDK lang. No attempt is made to distinguish between different SDK languages
	// here as processing a workflow with a different language than the one which authored it is
	// already undefined behavior. See `core_used_patches` for more.
	//
	// (-- api-linter: core::0141::forbidden-types=disabled
	//
	//	aip.dev/not-precedent: These really shouldn't have negative values. --)
	LangUsedFlags []uint32 `protobuf:"varint,2,rep,packed,name=lang_used_flags,json=langUsedFlags,proto3" json:"lang_used_flags,omitempty"`
	// Name of the SDK that processed the task. This is usually something like "temporal-go" and is
	// usually the same as client-name gRPC header. This should only be set if its value changed
	// since the last time recorded on the workflow (or be set on the first task).
	//
	// (-- api-linter: core::0122::name-suffix=disabled
	//
	//	aip.dev/not-precedent: We're ok with a name suffix here. --)
	SdkName string `protobuf:"bytes,3,opt,name=sdk_name,json=sdkName,proto3" json:"sdk_name,omitempty"`
	// Version of the SDK that processed the task. This is usually something like "1.20.0" and is
	// usually the same as client-version gRPC header. This should only be set if its value changed
	// since the last time recorded on the workflow (or be set on the first task).
	SdkVersion string `protobuf:"bytes,4,opt,name=sdk_version,json=sdkVersion,proto3" json:"sdk_version,omitempty"`
}

func (m *WorkflowTaskCompletedMetadata) Reset()      { *m = WorkflowTaskCompletedMetadata{} }
func (*WorkflowTaskCompletedMetadata) ProtoMessage() {}
func (*WorkflowTaskCompletedMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4c87c329f13a1874, []int{0}
}
func (m *WorkflowTaskCompletedMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkflowTaskCompletedMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkflowTaskCompletedMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkflowTaskCompletedMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowTaskCompletedMetadata.Merge(m, src)
}
func (m *WorkflowTaskCompletedMetadata) XXX_Size() int {
	return m.Size()
}
func (m *WorkflowTaskCompletedMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowTaskCompletedMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowTaskCompletedMetadata proto.InternalMessageInfo

func (m *WorkflowTaskCompletedMetadata) GetCoreUsedFlags() []uint32 {
	if m != nil {
		return m.CoreUsedFlags
	}
	return nil
}

func (m *WorkflowTaskCompletedMetadata) GetLangUsedFlags() []uint32 {
	if m != nil {
		return m.LangUsedFlags
	}
	return nil
}

func (m *WorkflowTaskCompletedMetadata) GetSdkName() string {
	if m != nil {
		return m.SdkName
	}
	return ""
}

func (m *WorkflowTaskCompletedMetadata) GetSdkVersion() string {
	if m != nil {
		return m.SdkVersion
	}
	return ""
}

func init() {
	proto.RegisterType((*WorkflowTaskCompletedMetadata)(nil), "temporal.api.sdk.v1.WorkflowTaskCompletedMetadata")
}

func init() {
	proto.RegisterFile("temporal/api/sdk/v1/task_complete_metadata.proto", fileDescriptor_4c87c329f13a1874)
}

var fileDescriptor_4c87c329f13a1874 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbf, 0x4e, 0x2a, 0x41,
	0x14, 0xc6, 0x77, 0xe0, 0xe6, 0xde, 0xeb, 0x18, 0x62, 0x82, 0xd1, 0x2c, 0x05, 0x47, 0x62, 0x61,
	0xa8, 0x66, 0xdd, 0xd8, 0xad, 0x15, 0x98, 0xd8, 0x69, 0x08, 0x20, 0x26, 0x36, 0x9b, 0x91, 0x19,
	0xc8, 0x64, 0x76, 0x99, 0xc9, 0xce, 0x8a, 0xad, 0x8f, 0xe0, 0x63, 0x18, 0x7d, 0x04, 0x5f, 0xc0,
	0x92, 0x92, 0x52, 0x86, 0xc6, 0x58, 0xf1, 0x08, 0x66, 0xf9, 0x13, 0xb6, 0xb0, 0xfd, 0x7d, 0xdf,
	0xaf, 0xf8, 0xce, 0xc1, 0xa7, 0x29, 0x8f, 0xb5, 0x4a, 0x68, 0xe4, 0x51, 0x2d, 0x3c, 0xc3, 0xa4,
	0x37, 0xf6, 0xbd, 0x94, 0x1a, 0x19, 0xf6, 0x55, 0xac, 0x23, 0x9e, 0xf2, 0x30, 0xe6, 0x29, 0x65,
	0x34, 0xa5, 0x44, 0x27, 0x2a, 0x55, 0xe5, 0xfd, 0x8d, 0x41, 0xa8, 0x16, 0xc4, 0x30, 0x49, 0xc6,
	0xfe, 0xf1, 0x1b, 0xc2, 0xd5, 0x5b, 0x95, 0xc8, 0x41, 0xa4, 0x1e, 0xbb, 0xd4, 0xc8, 0x8b, 0xb5,
	0xcc, 0xae, 0xd6, 0x72, 0xf9, 0x04, 0xef, 0xf5, 0x55, 0xc2, 0xc3, 0x07, 0xc3, 0x59, 0x38, 0x88,
	0xe8, 0xd0, 0xb8, 0xa8, 0x56, 0xac, 0x97, 0xda, 0xa5, 0x0c, 0xdf, 0x18, 0xce, 0x2e, 0x33, 0x98,
	0xf5, 0x22, 0x3a, 0x1a, 0xe6, 0x7b, 0x85, 0x55, 0x2f, 0xc3, 0xdb, 0x5e, 0x05, 0xff, 0x37, 0x4c,
	0x86, 0x23, 0x1a, 0x73, 0xb7, 0x58, 0x43, 0xf5, 0x9d, 0xf6, 0x3f, 0xc3, 0xe4, 0x35, 0x8d, 0x79,
	0xf9, 0x08, 0xef, 0x66, 0xd1, 0x98, 0x27, 0x46, 0xa8, 0x91, 0xfb, 0x67, 0x99, 0x62, 0xc3, 0x64,
	0x6f, 0x45, 0x9a, 0xef, 0x68, 0x32, 0x03, 0x67, 0x3a, 0x03, 0x67, 0x31, 0x03, 0xf4, 0x64, 0x01,
	0xbd, 0x58, 0x40, 0x1f, 0x16, 0xd0, 0xc4, 0x02, 0xfa, 0xb4, 0x80, 0xbe, 0x2c, 0x38, 0x0b, 0x0b,
	0xe8, 0x79, 0x0e, 0xce, 0x64, 0x0e, 0xce, 0x74, 0x0e, 0x0e, 0x3e, 0x14, 0x8a, 0xfc, 0x32, 0xbe,
	0x59, 0xc9, 0x2f, 0xde, 0x0c, 0x6e, 0x65, 0xc7, 0x6a, 0xa1, 0xbb, 0xea, 0x30, 0x27, 0x09, 0x95,
	0x3b, 0xf3, 0xb9, 0x61, 0xf2, 0xb5, 0x70, 0xd0, 0x5d, 0x87, 0x42, 0x91, 0x86, 0x16, 0xa4, 0xc3,
	0x24, 0xe9, 0xf9, 0xdf, 0x05, 0x77, 0xcb, 0x83, 0xa0, 0xa1, 0x45, 0x10, 0x74, 0x98, 0x0c, 0x82,
	0x9e, 0x7f, 0xff, 0x77, 0xf9, 0x87, 0xb3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x67, 0xe7,
	0x35, 0xbb, 0x01, 0x00, 0x00,
}

func (this *WorkflowTaskCompletedMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*WorkflowTaskCompletedMetadata)
	if !ok {
		that2, ok := that.(WorkflowTaskCompletedMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.CoreUsedFlags) != len(that1.CoreUsedFlags) {
		return false
	}
	for i := range this.CoreUsedFlags {
		if this.CoreUsedFlags[i] != that1.CoreUsedFlags[i] {
			return false
		}
	}
	if len(this.LangUsedFlags) != len(that1.LangUsedFlags) {
		return false
	}
	for i := range this.LangUsedFlags {
		if this.LangUsedFlags[i] != that1.LangUsedFlags[i] {
			return false
		}
	}
	if this.SdkName != that1.SdkName {
		return false
	}
	if this.SdkVersion != that1.SdkVersion {
		return false
	}
	return true
}
func (this *WorkflowTaskCompletedMetadata) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&sdk.WorkflowTaskCompletedMetadata{")
	s = append(s, "CoreUsedFlags: "+fmt.Sprintf("%#v", this.CoreUsedFlags)+",\n")
	s = append(s, "LangUsedFlags: "+fmt.Sprintf("%#v", this.LangUsedFlags)+",\n")
	s = append(s, "SdkName: "+fmt.Sprintf("%#v", this.SdkName)+",\n")
	s = append(s, "SdkVersion: "+fmt.Sprintf("%#v", this.SdkVersion)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringTaskCompleteMetadata(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *WorkflowTaskCompletedMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkflowTaskCompletedMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkflowTaskCompletedMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SdkVersion) > 0 {
		i -= len(m.SdkVersion)
		copy(dAtA[i:], m.SdkVersion)
		i = encodeVarintTaskCompleteMetadata(dAtA, i, uint64(len(m.SdkVersion)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SdkName) > 0 {
		i -= len(m.SdkName)
		copy(dAtA[i:], m.SdkName)
		i = encodeVarintTaskCompleteMetadata(dAtA, i, uint64(len(m.SdkName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.LangUsedFlags) > 0 {
		dAtA2 := make([]byte, len(m.LangUsedFlags)*10)
		var j1 int
		for _, num := range m.LangUsedFlags {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintTaskCompleteMetadata(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CoreUsedFlags) > 0 {
		dAtA4 := make([]byte, len(m.CoreUsedFlags)*10)
		var j3 int
		for _, num := range m.CoreUsedFlags {
			for num >= 1<<7 {
				dAtA4[j3] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j3++
			}
			dAtA4[j3] = uint8(num)
			j3++
		}
		i -= j3
		copy(dAtA[i:], dAtA4[:j3])
		i = encodeVarintTaskCompleteMetadata(dAtA, i, uint64(j3))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTaskCompleteMetadata(dAtA []byte, offset int, v uint64) int {
	offset -= sovTaskCompleteMetadata(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WorkflowTaskCompletedMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CoreUsedFlags) > 0 {
		l = 0
		for _, e := range m.CoreUsedFlags {
			l += sovTaskCompleteMetadata(uint64(e))
		}
		n += 1 + sovTaskCompleteMetadata(uint64(l)) + l
	}
	if len(m.LangUsedFlags) > 0 {
		l = 0
		for _, e := range m.LangUsedFlags {
			l += sovTaskCompleteMetadata(uint64(e))
		}
		n += 1 + sovTaskCompleteMetadata(uint64(l)) + l
	}
	l = len(m.SdkName)
	if l > 0 {
		n += 1 + l + sovTaskCompleteMetadata(uint64(l))
	}
	l = len(m.SdkVersion)
	if l > 0 {
		n += 1 + l + sovTaskCompleteMetadata(uint64(l))
	}
	return n
}

func sovTaskCompleteMetadata(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTaskCompleteMetadata(x uint64) (n int) {
	return sovTaskCompleteMetadata(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *WorkflowTaskCompletedMetadata) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&WorkflowTaskCompletedMetadata{`,
		`CoreUsedFlags:` + fmt.Sprintf("%v", this.CoreUsedFlags) + `,`,
		`LangUsedFlags:` + fmt.Sprintf("%v", this.LangUsedFlags) + `,`,
		`SdkName:` + fmt.Sprintf("%v", this.SdkName) + `,`,
		`SdkVersion:` + fmt.Sprintf("%v", this.SdkVersion) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringTaskCompleteMetadata(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *WorkflowTaskCompletedMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTaskCompleteMetadata
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WorkflowTaskCompletedMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkflowTaskCompletedMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTaskCompleteMetadata
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.CoreUsedFlags = append(m.CoreUsedFlags, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTaskCompleteMetadata
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTaskCompleteMetadata
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTaskCompleteMetadata
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.CoreUsedFlags) == 0 {
					m.CoreUsedFlags = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTaskCompleteMetadata
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.CoreUsedFlags = append(m.CoreUsedFlags, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field CoreUsedFlags", wireType)
			}
		case 2:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTaskCompleteMetadata
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.LangUsedFlags = append(m.LangUsedFlags, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTaskCompleteMetadata
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthTaskCompleteMetadata
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthTaskCompleteMetadata
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.LangUsedFlags) == 0 {
					m.LangUsedFlags = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTaskCompleteMetadata
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.LangUsedFlags = append(m.LangUsedFlags, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field LangUsedFlags", wireType)
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SdkName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskCompleteMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTaskCompleteMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaskCompleteMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SdkName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SdkVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTaskCompleteMetadata
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTaskCompleteMetadata
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTaskCompleteMetadata
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SdkVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTaskCompleteMetadata(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTaskCompleteMetadata
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTaskCompleteMetadata
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTaskCompleteMetadata(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTaskCompleteMetadata
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTaskCompleteMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTaskCompleteMetadata
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTaskCompleteMetadata
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTaskCompleteMetadata
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTaskCompleteMetadata
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTaskCompleteMetadata        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTaskCompleteMetadata          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTaskCompleteMetadata = fmt.Errorf("proto: unexpected end of group")
)

// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/retention/retention.proto

package retention

import (
	protoreflect "github.com/gucooing/zzz/protobuf/reflect/protoreflect"
	protoimpl "github.com/gucooing/zzz/protobuf/runtime/protoimpl"
	descriptorpb "github.com/gucooing/zzz/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

type TopLevelEnum int32

const (
	TopLevelEnum_TOP_LEVEL_UNKNOWN TopLevelEnum = 0
)

// Enum value maps for TopLevelEnum.
var (
	TopLevelEnum_name = map[int32]string{
		0: "TOP_LEVEL_UNKNOWN",
	}
	TopLevelEnum_value = map[string]int32{
		"TOP_LEVEL_UNKNOWN": 0,
	}
)

func (x TopLevelEnum) Enum() *TopLevelEnum {
	p := new(TopLevelEnum)
	*p = x
	return p
}

func (x TopLevelEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TopLevelEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_retention_retention_proto_enumTypes[0].Descriptor()
}

func (TopLevelEnum) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_retention_retention_proto_enumTypes[0]
}

func (x TopLevelEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TopLevelEnum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TopLevelEnum(num)
	return nil
}

// Deprecated: Use TopLevelEnum.Descriptor instead.
func (TopLevelEnum) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDescGZIP(), []int{0}
}

type TopLevelMessage_NestedEnum int32

const (
	TopLevelMessage_NESTED_UNKNOWN TopLevelMessage_NestedEnum = 0
)

// Enum value maps for TopLevelMessage_NestedEnum.
var (
	TopLevelMessage_NestedEnum_name = map[int32]string{
		0: "NESTED_UNKNOWN",
	}
	TopLevelMessage_NestedEnum_value = map[string]int32{
		"NESTED_UNKNOWN": 0,
	}
)

func (x TopLevelMessage_NestedEnum) Enum() *TopLevelMessage_NestedEnum {
	p := new(TopLevelMessage_NestedEnum)
	*p = x
	return p
}

func (x TopLevelMessage_NestedEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TopLevelMessage_NestedEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_cmd_protoc_gen_go_testdata_retention_retention_proto_enumTypes[1].Descriptor()
}

func (TopLevelMessage_NestedEnum) Type() protoreflect.EnumType {
	return &file_cmd_protoc_gen_go_testdata_retention_retention_proto_enumTypes[1]
}

func (x TopLevelMessage_NestedEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *TopLevelMessage_NestedEnum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = TopLevelMessage_NestedEnum(num)
	return nil
}

// Deprecated: Use TopLevelMessage_NestedEnum.Descriptor instead.
func (TopLevelMessage_NestedEnum) EnumDescriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDescGZIP(), []int{1, 0}
}

type Extendee struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	extensionFields protoimpl.ExtensionFields
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Extendee) Reset() {
	*x = Extendee{}
	mi := &file_cmd_protoc_gen_go_testdata_retention_retention_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Extendee) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extendee) ProtoMessage() {}

func (x *Extendee) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_retention_retention_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extendee.ProtoReflect.Descriptor instead.
func (*Extendee) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDescGZIP(), []int{0}
}

type TopLevelMessage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	F     *float32               `protobuf:"fixed32,1,opt,name=f" json:"f,omitempty"`
	// Types that are valid to be assigned to O:
	//
	//	*TopLevelMessage_I
	O               isTopLevelMessage_O `protobuf_oneof:"o"`
	extensionFields protoimpl.ExtensionFields
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *TopLevelMessage) Reset() {
	*x = TopLevelMessage{}
	mi := &file_cmd_protoc_gen_go_testdata_retention_retention_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TopLevelMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopLevelMessage) ProtoMessage() {}

func (x *TopLevelMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_retention_retention_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopLevelMessage.ProtoReflect.Descriptor instead.
func (*TopLevelMessage) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDescGZIP(), []int{1}
}

func (x *TopLevelMessage) GetF() float32 {
	if x != nil && x.F != nil {
		return *x.F
	}
	return 0
}

func (x *TopLevelMessage) GetO() isTopLevelMessage_O {
	if x != nil {
		return x.O
	}
	return nil
}

func (x *TopLevelMessage) GetI() int64 {
	if x != nil {
		if x, ok := x.O.(*TopLevelMessage_I); ok {
			return x.I
		}
	}
	return 0
}

type isTopLevelMessage_O interface {
	isTopLevelMessage_O()
}

type TopLevelMessage_I struct {
	I int64 `protobuf:"varint,2,opt,name=i,oneof"`
}

func (*TopLevelMessage_I) isTopLevelMessage_O() {}

type TopLevelMessage_NestedMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TopLevelMessage_NestedMessage) Reset() {
	*x = TopLevelMessage_NestedMessage{}
	mi := &file_cmd_protoc_gen_go_testdata_retention_retention_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TopLevelMessage_NestedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopLevelMessage_NestedMessage) ProtoMessage() {}

func (x *TopLevelMessage_NestedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_retention_retention_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopLevelMessage_NestedMessage.ProtoReflect.Descriptor instead.
func (*TopLevelMessage_NestedMessage) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDescGZIP(), []int{1, 0}
}

var file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         505092806,
		Name:          "testretention.plain_option",
		Tag:           "varint,505092806,opt,name=plain_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         505039132,
		Name:          "testretention.runtime_retention_option",
		Tag:           "varint,505039132,opt,name=runtime_retention_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         504878676,
		Name:          "testretention.source_retention_option",
		Tag:           "varint,504878676,opt,name=source_retention_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: ([]*OptionsMessage)(nil),
		Field:         504823570,
		Name:          "testretention.repeated_options",
		Tag:           "bytes,504823570,rep,name=repeated_options",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ExtensionRangeOptions)(nil),
		ExtensionType: (*OptionsMessage)(nil),
		Field:         504822148,
		Name:          "testretention.extension_range_option",
		Tag:           "bytes,504822148,opt,name=extension_range_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*OptionsMessage)(nil),
		Field:         504820819,
		Name:          "testretention.message_option",
		Tag:           "bytes,504820819,opt,name=message_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*OptionsMessage)(nil),
		Field:         504589219,
		Name:          "testretention.field_option",
		Tag:           "bytes,504589219,opt,name=field_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
	{
		ExtendedType:  (*descriptorpb.OneofOptions)(nil),
		ExtensionType: (*OptionsMessage)(nil),
		Field:         504479153,
		Name:          "testretention.oneof_option",
		Tag:           "bytes,504479153,opt,name=oneof_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*OptionsMessage)(nil),
		Field:         504451567,
		Name:          "testretention.enum_option",
		Tag:           "bytes,504451567,opt,name=enum_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*OptionsMessage)(nil),
		Field:         504450522,
		Name:          "testretention.enum_entry_option",
		Tag:           "bytes,504450522,opt,name=enum_entry_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
	{
		ExtendedType:  (*descriptorpb.ServiceOptions)(nil),
		ExtensionType: (*OptionsMessage)(nil),
		Field:         504387709,
		Name:          "testretention.service_option",
		Tag:           "bytes,504387709,opt,name=service_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*OptionsMessage)(nil),
		Field:         504349420,
		Name:          "testretention.method_option",
		Tag:           "bytes,504349420,opt,name=method_option",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
	{
		ExtendedType:  (*Extendee)(nil),
		ExtensionType: (*int32)(nil),
		Field:         1,
		Name:          "testretention.i",
		Tag:           "varint,1,opt,name=i",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
	{
		ExtendedType:  (*Extendee)(nil),
		ExtensionType: (*string)(nil),
		Field:         2,
		Name:          "testretention.TopLevelMessage.s",
		Tag:           "bytes,2,opt,name=s",
		Filename:      "cmd/protoc-gen-go/testdata/retention/retention.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional int32 plain_option = 505092806;
	E_PlainOption = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[0]
	// optional int32 runtime_retention_option = 505039132;
	E_RuntimeRetentionOption = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[1]
	// optional int32 source_retention_option = 504878676;
	E_SourceRetentionOption = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[2]
	// repeated testretention.OptionsMessage repeated_options = 504823570;
	E_RepeatedOptions = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[3]
)

// Extension fields to descriptorpb.ExtensionRangeOptions.
var (
	// optional testretention.OptionsMessage extension_range_option = 504822148;
	E_ExtensionRangeOption = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[4]
)

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional testretention.OptionsMessage message_option = 504820819;
	E_MessageOption = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[5]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional testretention.OptionsMessage field_option = 504589219;
	E_FieldOption = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[6]
)

// Extension fields to descriptorpb.OneofOptions.
var (
	// optional testretention.OptionsMessage oneof_option = 504479153;
	E_OneofOption = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[7]
)

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional testretention.OptionsMessage enum_option = 504451567;
	E_EnumOption = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[8]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional testretention.OptionsMessage enum_entry_option = 504450522;
	E_EnumEntryOption = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[9]
)

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional testretention.OptionsMessage service_option = 504387709;
	E_ServiceOption = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[10]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional testretention.OptionsMessage method_option = 504349420;
	E_MethodOption = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[11]
)

// Extension fields to Extendee.
var (
	// optional int32 i = 1;
	E_I = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[12]
	// optional string s = 2;
	E_TopLevelMessage_S = &file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes[13]
)

var File_cmd_protoc_gen_go_testdata_retention_retention_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDesc = []byte{
	0x0a, 0x34, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x72, 0x65, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x72, 0x65, 0x74, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3a, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x16, 0x0a, 0x08, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x65, 0x2a,
	0x04, 0x08, 0x01, 0x10, 0x02, 0x2a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0xe8, 0x01, 0x0a, 0x0f,
	0x54, 0x6f, 0x70, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x01, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0a, 0x9a, 0xba, 0xed, 0x84,
	0x0f, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x01, 0x66, 0x12, 0x0e, 0x0a, 0x01, 0x69, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x01, 0x69, 0x1a, 0x1b, 0x0a, 0x0d, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x3a, 0x0a, 0x9a, 0xc5, 0xde, 0x85,
	0x0f, 0x04, 0x08, 0x01, 0x10, 0x02, 0x22, 0x2c, 0x0a, 0x0a, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x45, 0x6e, 0x75, 0x6d, 0x12, 0x12, 0x0a, 0x0e, 0x4e, 0x45, 0x53, 0x54, 0x45, 0x44, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a, 0x0a, 0xfa, 0x9e, 0xaa, 0x84, 0x0f, 0x04,
	0x08, 0x01, 0x10, 0x02, 0x2a, 0x10, 0x08, 0x0a, 0x10, 0x65, 0x1a, 0x0a, 0xa2, 0x98, 0xdf, 0x85,
	0x0f, 0x04, 0x08, 0x01, 0x10, 0x02, 0x32, 0x31, 0x0a, 0x01, 0x73, 0x12, 0x17, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x9a, 0xba, 0xed, 0x84,
	0x0f, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x01, 0x73, 0x3a, 0x0a, 0x9a, 0xc5, 0xde, 0x85, 0x0f,
	0x04, 0x08, 0x01, 0x10, 0x02, 0x42, 0x0f, 0x0a, 0x01, 0x6f, 0x12, 0x0a, 0x8a, 0xdb, 0xb7, 0x84,
	0x0f, 0x04, 0x08, 0x01, 0x10, 0x02, 0x2a, 0x3d, 0x0a, 0x0c, 0x54, 0x6f, 0x70, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x11, 0x54, 0x4f, 0x50, 0x5f, 0x4c, 0x45,
	0x56, 0x45, 0x4c, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x1a, 0x0a, 0xd2,
	0xdd, 0xa9, 0x84, 0x0f, 0x04, 0x08, 0x01, 0x10, 0x02, 0x1a, 0x0a, 0xfa, 0x9e, 0xaa, 0x84, 0x0f,
	0x04, 0x08, 0x01, 0x10, 0x02, 0x32, 0x6c, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x55, 0x0a, 0x07, 0x44, 0x6f, 0x53, 0x74, 0x75, 0x66, 0x66, 0x12, 0x1e, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x70, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1e, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x6f, 0x70, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0a, 0xe2, 0xae, 0xf8,
	0x83, 0x0f, 0x04, 0x08, 0x01, 0x10, 0x02, 0x1a, 0x0a, 0xea, 0x87, 0x8b, 0x84, 0x0f, 0x04, 0x08,
	0x01, 0x10, 0x02, 0x3a, 0x43, 0x0a, 0x0c, 0x70, 0x6c, 0x61, 0x69, 0x6e, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xc6, 0xb5, 0xec, 0xf0, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x70, 0x6c, 0x61,
	0x69, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x5f, 0x0a, 0x18, 0x72, 0x75, 0x6e, 0x74,
	0x69, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x9c, 0x92, 0xe9, 0xf0, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0x88, 0x01,
	0x01, 0x52, 0x16, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x5d, 0x0a, 0x17, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xd4, 0xac, 0xdf, 0xf0, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x03, 0x88, 0x01,
	0x02, 0x52, 0x15, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x6a, 0x0a, 0x10, 0x72, 0x65, 0x70, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x92, 0xfe, 0xdb, 0xf0, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x65, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x65, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x3a, 0x7f, 0x0a, 0x16, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x84, 0xf3, 0xdb, 0xf0, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x14, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x69, 0x0a, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd3, 0xe8, 0xdb, 0xf0, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x3a, 0x63, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xa3, 0xd7, 0xcd, 0xf0, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x63, 0x0a, 0x0c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xb1, 0xfb, 0xc6, 0xf0, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x60, 0x0a, 0x0b, 0x65, 0x6e,
	0x75, 0x6d, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xef, 0xa3, 0xc5, 0xf0, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x70, 0x0a, 0x11,
	0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xda, 0x9b, 0xc5, 0xf0, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0f, 0x65,
	0x6e, 0x75, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x69,
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xfd, 0xb0, 0xc1, 0xf0, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x66, 0x0a, 0x0d, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xec, 0x85, 0xbf, 0xf0, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x65, 0x74, 0x65, 0x6e,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x3a, 0x31, 0x0a, 0x01, 0x69, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x65, 0x74,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x9a, 0xba, 0xed, 0x84, 0x0f, 0x04, 0x08, 0x01, 0x10,
	0x02, 0x52, 0x01, 0x69, 0x42, 0x6d, 0x92, 0xf1, 0xdf, 0x85, 0x0f, 0x04, 0x08, 0x01, 0x10, 0x02,
	0x82, 0x90, 0xf7, 0x85, 0x0f, 0x04, 0x08, 0x01, 0x10, 0x02, 0xe0, 0x91, 0xc9, 0x86, 0x0f, 0x02,
	0xb0, 0xac, 0xe3, 0x86, 0x0f, 0x01, 0x80, 0x9b, 0xb2, 0xa0, 0x0f, 0x01, 0xd0, 0xae, 0xfc, 0xa2,
	0x0f, 0x02, 0x5a, 0x3f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63,
	0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e,
}

var (
	file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDescData = file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_retention_retention_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cmd_protoc_gen_go_testdata_retention_retention_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cmd_protoc_gen_go_testdata_retention_retention_proto_goTypes = []any{
	(TopLevelEnum)(0),                          // 0: testretention.TopLevelEnum
	(TopLevelMessage_NestedEnum)(0),            // 1: testretention.TopLevelMessage.NestedEnum
	(*Extendee)(nil),                           // 2: testretention.Extendee
	(*TopLevelMessage)(nil),                    // 3: testretention.TopLevelMessage
	(*TopLevelMessage_NestedMessage)(nil),      // 4: testretention.TopLevelMessage.NestedMessage
	(*descriptorpb.FileOptions)(nil),           // 5: google.protobuf.FileOptions
	(*descriptorpb.ExtensionRangeOptions)(nil), // 6: google.protobuf.ExtensionRangeOptions
	(*descriptorpb.MessageOptions)(nil),        // 7: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),          // 8: google.protobuf.FieldOptions
	(*descriptorpb.OneofOptions)(nil),          // 9: google.protobuf.OneofOptions
	(*descriptorpb.EnumOptions)(nil),           // 10: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil),      // 11: google.protobuf.EnumValueOptions
	(*descriptorpb.ServiceOptions)(nil),        // 12: google.protobuf.ServiceOptions
	(*descriptorpb.MethodOptions)(nil),         // 13: google.protobuf.MethodOptions
	(*OptionsMessage)(nil),                     // 14: testretention.OptionsMessage
}
var file_cmd_protoc_gen_go_testdata_retention_retention_proto_depIdxs = []int32{
	5,  // 0: testretention.plain_option:extendee -> google.protobuf.FileOptions
	5,  // 1: testretention.runtime_retention_option:extendee -> google.protobuf.FileOptions
	5,  // 2: testretention.source_retention_option:extendee -> google.protobuf.FileOptions
	5,  // 3: testretention.repeated_options:extendee -> google.protobuf.FileOptions
	6,  // 4: testretention.extension_range_option:extendee -> google.protobuf.ExtensionRangeOptions
	7,  // 5: testretention.message_option:extendee -> google.protobuf.MessageOptions
	8,  // 6: testretention.field_option:extendee -> google.protobuf.FieldOptions
	9,  // 7: testretention.oneof_option:extendee -> google.protobuf.OneofOptions
	10, // 8: testretention.enum_option:extendee -> google.protobuf.EnumOptions
	11, // 9: testretention.enum_entry_option:extendee -> google.protobuf.EnumValueOptions
	12, // 10: testretention.service_option:extendee -> google.protobuf.ServiceOptions
	13, // 11: testretention.method_option:extendee -> google.protobuf.MethodOptions
	2,  // 12: testretention.i:extendee -> testretention.Extendee
	2,  // 13: testretention.TopLevelMessage.s:extendee -> testretention.Extendee
	14, // 14: testretention.repeated_options:type_name -> testretention.OptionsMessage
	14, // 15: testretention.extension_range_option:type_name -> testretention.OptionsMessage
	14, // 16: testretention.message_option:type_name -> testretention.OptionsMessage
	14, // 17: testretention.field_option:type_name -> testretention.OptionsMessage
	14, // 18: testretention.oneof_option:type_name -> testretention.OptionsMessage
	14, // 19: testretention.enum_option:type_name -> testretention.OptionsMessage
	14, // 20: testretention.enum_entry_option:type_name -> testretention.OptionsMessage
	14, // 21: testretention.service_option:type_name -> testretention.OptionsMessage
	14, // 22: testretention.method_option:type_name -> testretention.OptionsMessage
	3,  // 23: testretention.Service.DoStuff:input_type -> testretention.TopLevelMessage
	3,  // 24: testretention.Service.DoStuff:output_type -> testretention.TopLevelMessage
	24, // [24:25] is the sub-list for method output_type
	23, // [23:24] is the sub-list for method input_type
	14, // [14:23] is the sub-list for extension type_name
	0,  // [0:14] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_retention_retention_proto_init() }
func file_cmd_protoc_gen_go_testdata_retention_retention_proto_init() {
	if File_cmd_protoc_gen_go_testdata_retention_retention_proto != nil {
		return
	}
	file_cmd_protoc_gen_go_testdata_retention_options_message_proto_init()
	file_cmd_protoc_gen_go_testdata_retention_retention_proto_msgTypes[1].OneofWrappers = []any{
		(*TopLevelMessage_I)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 14,
			NumServices:   1,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_retention_retention_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_retention_retention_proto_depIdxs,
		EnumInfos:         file_cmd_protoc_gen_go_testdata_retention_retention_proto_enumTypes,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_retention_retention_proto_msgTypes,
		ExtensionInfos:    file_cmd_protoc_gen_go_testdata_retention_retention_proto_extTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_retention_retention_proto = out.File
	file_cmd_protoc_gen_go_testdata_retention_retention_proto_rawDesc = nil
	file_cmd_protoc_gen_go_testdata_retention_retention_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_retention_retention_proto_depIdxs = nil
}

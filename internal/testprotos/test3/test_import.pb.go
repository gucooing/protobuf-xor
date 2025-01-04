// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/test3/test_import.proto

package test3

import (
	protoreflect protobuf "github.com/gucooing/protobuf-xor/reflect/protoreflect"
	protoimpl protobuf "github.com/gucooing/protobuf-xor/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type ImportEnum int32

const (
	ImportEnum_IMPORT_ZERO ImportEnum = 0
)

// Enum value maps for ImportEnum.
var (
	ImportEnum_name = map[int32]string{
		0: "IMPORT_ZERO",
	}
	ImportEnum_value = map[string]int32{
		"IMPORT_ZERO": 0,
	}
)

func (x ImportEnum) Enum() *ImportEnum {
	p := new(ImportEnum)
	*p = x
	return p
}

func (x ImportEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImportEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_test3_test_import_proto_enumTypes[0].Descriptor()
}

func (ImportEnum) Type() protoreflect.EnumType {
	return &file_internal_testprotos_test3_test_import_proto_enumTypes[0]
}

func (x ImportEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImportEnum.Descriptor instead.
func (ImportEnum) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_test3_test_import_proto_rawDescGZIP(), []int{0}
}

type ImportMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImportMessage) Reset() {
	*x = ImportMessage{}
	mi := &file_internal_testprotos_test3_test_import_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImportMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportMessage) ProtoMessage() {}

func (x *ImportMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_test3_test_import_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportMessage.ProtoReflect.Descriptor instead.
func (*ImportMessage) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_test3_test_import_proto_rawDescGZIP(), []int{0}
}

var File_internal_testprotos_test3_test_import_proto protoreflect.FileDescriptor

var file_internal_testprotos_test3_test_import_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x33, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x33, 0x22, 0x0f, 0x0a, 0x0d, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2a, 0x1d, 0x0a, 0x0a, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x45, 0x6e, 0x75,
	0x6d, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4d, 0x50, 0x4f, 0x52, 0x54, 0x5f, 0x5a, 0x45, 0x52, 0x4f,
	0x10, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_testprotos_test3_test_import_proto_rawDescOnce sync.Once
	file_internal_testprotos_test3_test_import_proto_rawDescData = file_internal_testprotos_test3_test_import_proto_rawDesc
)

func file_internal_testprotos_test3_test_import_proto_rawDescGZIP() []byte {
	file_internal_testprotos_test3_test_import_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_test3_test_import_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_test3_test_import_proto_rawDescData)
	})
	return file_internal_testprotos_test3_test_import_proto_rawDescData
}

var file_internal_testprotos_test3_test_import_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_testprotos_test3_test_import_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_test3_test_import_proto_goTypes = []any{
	(ImportEnum)(0),       // 0: goproto.proto.test3.ImportEnum
	(*ImportMessage)(nil), // 1: goproto.proto.test3.ImportMessage
}
var file_internal_testprotos_test3_test_import_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_test3_test_import_proto_init() }
func file_internal_testprotos_test3_test_import_proto_init() {
	if File_internal_testprotos_test3_test_import_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_testprotos_test3_test_import_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_test3_test_import_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_test3_test_import_proto_depIdxs,
		EnumInfos:         file_internal_testprotos_test3_test_import_proto_enumTypes,
		MessageInfos:      file_internal_testprotos_test3_test_import_proto_msgTypes,
	}.Build()
	File_internal_testprotos_test3_test_import_proto = out.File
	file_internal_testprotos_test3_test_import_proto_rawDesc = nil
	file_internal_testprotos_test3_test_import_proto_goTypes = nil
	file_internal_testprotos_test3_test_import_proto_depIdxs = nil
}

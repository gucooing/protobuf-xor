// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/race/extender/test.proto

package extender

import (
	message "github.com/gucooing/protobuf-xor/internal/testprotos/race/message"
	protoreflect "github.com/gucooing/protobuf-xor/reflect/protoreflect"
	protoimpl "github.com/gucooing/protobuf-xor/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type OtherMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	I32           *int32                 `protobuf:"varint,1,opt,name=i32" json:"i32,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OtherMessage) Reset() {
	*x = OtherMessage{}
	mi := &file_internal_testprotos_race_extender_test_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OtherMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OtherMessage) ProtoMessage() {}

func (x *OtherMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_race_extender_test_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OtherMessage.ProtoReflect.Descriptor instead.
func (*OtherMessage) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_race_extender_test_proto_rawDescGZIP(), []int{0}
}

func (x *OtherMessage) GetI32() int32 {
	if x != nil && x.I32 != nil {
		return *x.I32
	}
	return 0
}

var file_internal_testprotos_race_extender_test_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*message.MyMessage)(nil),
		ExtensionType: (*string)(nil),
		Field:         2,
		Name:          "goproto.proto.test.s",
		Tag:           "bytes,2,opt,name=s",
		Filename:      "internal/testprotos/race/extender/test.proto",
	},
}

// Extension fields to message.MyMessage.
var (
	// optional string s = 2;
	E_S = &file_internal_testprotos_race_extender_test_proto_extTypes[0]
)

var File_internal_testprotos_race_extender_test_proto protoreflect.FileDescriptor

var file_internal_testprotos_race_extender_test_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x20, 0x0a, 0x0c, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x69, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x33,
	0x32, 0x3a, 0x2b, 0x0a, 0x01, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x73, 0x42, 0x3e,
	0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x72, 0x61, 0x63, 0x65, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x62, 0x08,
	0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
}

var (
	file_internal_testprotos_race_extender_test_proto_rawDescOnce sync.Once
	file_internal_testprotos_race_extender_test_proto_rawDescData = file_internal_testprotos_race_extender_test_proto_rawDesc
)

func file_internal_testprotos_race_extender_test_proto_rawDescGZIP() []byte {
	file_internal_testprotos_race_extender_test_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_race_extender_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_race_extender_test_proto_rawDescData)
	})
	return file_internal_testprotos_race_extender_test_proto_rawDescData
}

var file_internal_testprotos_race_extender_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_testprotos_race_extender_test_proto_goTypes = []any{
	(*OtherMessage)(nil),      // 0: goproto.proto.test.OtherMessage
	(*message.MyMessage)(nil), // 1: goproto.proto.test.MyMessage
}
var file_internal_testprotos_race_extender_test_proto_depIdxs = []int32{
	1, // 0: goproto.proto.test.s:extendee -> goproto.proto.test.MyMessage
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_race_extender_test_proto_init() }
func file_internal_testprotos_race_extender_test_proto_init() {
	if File_internal_testprotos_race_extender_test_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_testprotos_race_extender_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_race_extender_test_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_race_extender_test_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_race_extender_test_proto_msgTypes,
		ExtensionInfos:    file_internal_testprotos_race_extender_test_proto_extTypes,
	}.Build()
	File_internal_testprotos_race_extender_test_proto = out.File
	file_internal_testprotos_race_extender_test_proto_rawDesc = nil
	file_internal_testprotos_race_extender_test_proto_goTypes = nil
	file_internal_testprotos_race_extender_test_proto_depIdxs = nil
}

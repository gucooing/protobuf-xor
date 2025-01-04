// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/messageset/msetextpb/msetextpb.proto

package msetextpb

import (
	messagesetpb protobuf "github.com/gucooing/protobuf-xor/internal/testprotos/messageset/messagesetpb"
	protoreflect protobuf "github.com/gucooing/protobuf-xor/reflect/protoreflect"
	protoimpl protobuf "github.com/gucooing/protobuf-xor/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Ext1 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ext1Field1    *int32                 `protobuf:"varint,1,opt,name=ext1_field1,json=ext1Field1" json:"ext1_field1,omitempty"`
	Ext1Field2    *int32                 `protobuf:"varint,2,opt,name=ext1_field2,json=ext1Field2" json:"ext1_field2,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ext1) Reset() {
	*x = Ext1{}
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ext1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ext1) ProtoMessage() {}

func (x *Ext1) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ext1.ProtoReflect.Descriptor instead.
func (*Ext1) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDescGZIP(), []int{0}
}

func (x *Ext1) GetExt1Field1() int32 {
	if x != nil && x.Ext1Field1 != nil {
		return *x.Ext1Field1
	}
	return 0
}

func (x *Ext1) GetExt1Field2() int32 {
	if x != nil && x.Ext1Field2 != nil {
		return *x.Ext1Field2
	}
	return 0
}

type Ext2 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ext2Field1    *int32                 `protobuf:"varint,1,opt,name=ext2_field1,json=ext2Field1" json:"ext2_field1,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Ext2) Reset() {
	*x = Ext2{}
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Ext2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ext2) ProtoMessage() {}

func (x *Ext2) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ext2.ProtoReflect.Descriptor instead.
func (*Ext2) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDescGZIP(), []int{1}
}

func (x *Ext2) GetExt2Field1() int32 {
	if x != nil && x.Ext2Field1 != nil {
		return *x.Ext2Field1
	}
	return 0
}

type ExtRequired struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	RequiredField1 *int32                 `protobuf:"varint,1,req,name=required_field1,json=requiredField1" json:"required_field1,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ExtRequired) Reset() {
	*x = ExtRequired{}
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtRequired) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtRequired) ProtoMessage() {}

func (x *ExtRequired) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtRequired.ProtoReflect.Descriptor instead.
func (*ExtRequired) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDescGZIP(), []int{2}
}

func (x *ExtRequired) GetRequiredField1() int32 {
	if x != nil && x.RequiredField1 != nil {
		return *x.RequiredField1
	}
	return 0
}

type ExtLargeNumber struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExtLargeNumber) Reset() {
	*x = ExtLargeNumber{}
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtLargeNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtLargeNumber) ProtoMessage() {}

func (x *ExtLargeNumber) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_messageset_msetextpb_msetextpb_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtLargeNumber.ProtoReflect.Descriptor instead.
func (*ExtLargeNumber) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDescGZIP(), []int{3}
}

var file_internal_testprotos_messageset_msetextpb_msetextpb_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*messagesetpb.MessageSet)(nil),
		ExtensionType: (*Ext1)(nil),
		Field:         1000,
		Name:          "goproto.proto.messageset.Ext1.message_set_ext1",
		Tag:           "bytes,1000,opt,name=message_set_ext1",
		Filename:      "internal/testprotos/messageset/msetextpb/msetextpb.proto",
	},
	{
		ExtendedType:  (*messagesetpb.MessageSet)(nil),
		ExtensionType: (*Ext2)(nil),
		Field:         1001,
		Name:          "goproto.proto.messageset.Ext2.message_set_ext2",
		Tag:           "bytes,1001,opt,name=message_set_ext2",
		Filename:      "internal/testprotos/messageset/msetextpb/msetextpb.proto",
	},
	{
		ExtendedType:  (*messagesetpb.MessageSet)(nil),
		ExtensionType: (*ExtRequired)(nil),
		Field:         1002,
		Name:          "goproto.proto.messageset.ExtRequired.message_set_extrequired",
		Tag:           "bytes,1002,opt,name=message_set_extrequired",
		Filename:      "internal/testprotos/messageset/msetextpb/msetextpb.proto",
	},
	{
		ExtendedType:  (*messagesetpb.MessageSet)(nil),
		ExtensionType: (*ExtLargeNumber)(nil),
		Field:         536870912,
		Name:          "goproto.proto.messageset.ExtLargeNumber.message_set_extlarge",
		Tag:           "bytes,536870912,opt,name=message_set_extlarge",
		Filename:      "internal/testprotos/messageset/msetextpb/msetextpb.proto",
	},
}

// Extension fields to messagesetpb.MessageSet.
var (
	// optional goproto.proto.messageset.Ext1 message_set_ext1 = 1000;
	E_Ext1_MessageSetExt1 = &file_internal_testprotos_messageset_msetextpb_msetextpb_proto_extTypes[0]
	// optional goproto.proto.messageset.Ext2 message_set_ext2 = 1001;
	E_Ext2_MessageSetExt2 = &file_internal_testprotos_messageset_msetextpb_msetextpb_proto_extTypes[1]
	// optional goproto.proto.messageset.ExtRequired message_set_extrequired = 1002;
	E_ExtRequired_MessageSetExtrequired = &file_internal_testprotos_messageset_msetextpb_msetextpb_proto_extTypes[2]
	// optional goproto.proto.messageset.ExtLargeNumber message_set_extlarge = 536870912;
	E_ExtLargeNumber_MessageSetExtlarge = &file_internal_testprotos_messageset_msetextpb_msetextpb_proto_extTypes[3] // 1<<29
)

var File_internal_testprotos_messageset_msetextpb_msetextpb_proto protoreflect.FileDescriptor

var file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDesc = []byte{
	0x0a, 0x38, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74,
	0x2f, 0x6d, 0x73, 0x65, 0x74, 0x65, 0x78, 0x74, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x65, 0x74, 0x65,
	0x78, 0x74, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x65, 0x74, 0x1a, 0x3d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x65, 0x74, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x70,
	0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x01, 0x0a, 0x04, 0x45, 0x78, 0x74, 0x31, 0x12, 0x1f, 0x0a, 0x0b,
	0x65, 0x78, 0x74, 0x31, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x31, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x12, 0x1f, 0x0a,
	0x0b, 0x65, 0x78, 0x74, 0x31, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x31, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x32, 0x6f,
	0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x78,
	0x74, 0x31, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x31, 0x52,
	0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74, 0x31, 0x22,
	0x98, 0x01, 0x0a, 0x04, 0x45, 0x78, 0x74, 0x32, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x74, 0x32,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x65,
	0x78, 0x74, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x6f, 0x0a, 0x10, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x32, 0x12, 0x24, 0x2e,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x32, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74, 0x32, 0x22, 0xc4, 0x01, 0x0a, 0x0b, 0x45,
	0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x0f, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x31, 0x32, 0x84, 0x01, 0x0a, 0x17, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x24, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65,
	0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x18, 0xea, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x45,
	0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x52, 0x15, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x45, 0x78, 0x74, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x22, 0x97, 0x01, 0x0a, 0x0e, 0x45, 0x78, 0x74, 0x4c, 0x61, 0x72, 0x67, 0x65, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x32, 0x84, 0x01, 0x0a, 0x14, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x12, 0x24, 0x2e,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x18, 0x80, 0x80, 0x80, 0x80, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x45, 0x78, 0x74, 0x4c, 0x61, 0x72, 0x67,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x45, 0x78, 0x74, 0x6c, 0x61, 0x72, 0x67, 0x65, 0x42, 0x45, 0x5a, 0x43, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x65, 0x74, 0x2f, 0x6d, 0x73, 0x65, 0x74, 0x65, 0x78, 0x74,
	0x70, 0x62, 0x62, 0x08, 0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
}

var (
	file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDescOnce sync.Once
	file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDescData = file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDesc
)

func file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDescGZIP() []byte {
	file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDescData)
	})
	return file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDescData
}

var file_internal_testprotos_messageset_msetextpb_msetextpb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_testprotos_messageset_msetextpb_msetextpb_proto_goTypes = []any{
	(*Ext1)(nil),                    // 0: goproto.proto.messageset.Ext1
	(*Ext2)(nil),                    // 1: goproto.proto.messageset.Ext2
	(*ExtRequired)(nil),             // 2: goproto.proto.messageset.ExtRequired
	(*ExtLargeNumber)(nil),          // 3: goproto.proto.messageset.ExtLargeNumber
	(*messagesetpb.MessageSet)(nil), // 4: goproto.proto.messageset.MessageSet
}
var file_internal_testprotos_messageset_msetextpb_msetextpb_proto_depIdxs = []int32{
	4, // 0: goproto.proto.messageset.Ext1.message_set_ext1:extendee -> goproto.proto.messageset.MessageSet
	4, // 1: goproto.proto.messageset.Ext2.message_set_ext2:extendee -> goproto.proto.messageset.MessageSet
	4, // 2: goproto.proto.messageset.ExtRequired.message_set_extrequired:extendee -> goproto.proto.messageset.MessageSet
	4, // 3: goproto.proto.messageset.ExtLargeNumber.message_set_extlarge:extendee -> goproto.proto.messageset.MessageSet
	0, // 4: goproto.proto.messageset.Ext1.message_set_ext1:type_name -> goproto.proto.messageset.Ext1
	1, // 5: goproto.proto.messageset.Ext2.message_set_ext2:type_name -> goproto.proto.messageset.Ext2
	2, // 6: goproto.proto.messageset.ExtRequired.message_set_extrequired:type_name -> goproto.proto.messageset.ExtRequired
	3, // 7: goproto.proto.messageset.ExtLargeNumber.message_set_extlarge:type_name -> goproto.proto.messageset.ExtLargeNumber
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	4, // [4:8] is the sub-list for extension type_name
	0, // [0:4] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_testprotos_messageset_msetextpb_msetextpb_proto_init() }
func file_internal_testprotos_messageset_msetextpb_msetextpb_proto_init() {
	if File_internal_testprotos_messageset_msetextpb_msetextpb_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_messageset_msetextpb_msetextpb_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_messageset_msetextpb_msetextpb_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_messageset_msetextpb_msetextpb_proto_msgTypes,
		ExtensionInfos:    file_internal_testprotos_messageset_msetextpb_msetextpb_proto_extTypes,
	}.Build()
	File_internal_testprotos_messageset_msetextpb_msetextpb_proto = out.File
	file_internal_testprotos_messageset_msetextpb_msetextpb_proto_rawDesc = nil
	file_internal_testprotos_messageset_msetextpb_msetextpb_proto_goTypes = nil
	file_internal_testprotos_messageset_msetextpb_msetextpb_proto_depIdxs = nil
}

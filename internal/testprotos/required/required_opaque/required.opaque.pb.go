// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/required/required_opaque/required.opaque.proto

package required_opaque

import (
	protoreflect protobuf "github.com/gucooing/protobuf-xor/reflect/protoreflect"
	protoimpl protobuf "github.com/gucooing/protobuf-xor/runtime/protoimpl"
	_ protobuf "github.com/gucooing/protobuf-xor/types/gofeaturespb"
	reflect "reflect"
)

type Int32 struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           int32                  `protobuf:"varint,1,req,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Int32) Reset() {
	*x = Int32{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Int32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int32) ProtoMessage() {}

func (x *Int32) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Int32) GetV() int32 {
	if x != nil {
		return x.xxx_hidden_V
	}
	return 0
}

func (x *Int32) SetV(v int32) {
	x.xxx_hidden_V = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Int32) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Int32) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = 0
}

type Int32_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *int32
}

func (b0 Int32_builder) Build() *Int32 {
	m0 := &Int32{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = *b.V
	}
	return m0
}

type Int64 struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           int64                  `protobuf:"varint,1,req,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Int64) Reset() {
	*x = Int64{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Int64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Int64) ProtoMessage() {}

func (x *Int64) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Int64) GetV() int64 {
	if x != nil {
		return x.xxx_hidden_V
	}
	return 0
}

func (x *Int64) SetV(v int64) {
	x.xxx_hidden_V = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Int64) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Int64) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = 0
}

type Int64_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *int64
}

func (b0 Int64_builder) Build() *Int64 {
	m0 := &Int64{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = *b.V
	}
	return m0
}

type Uint32 struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           uint32                 `protobuf:"varint,1,req,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Uint32) Reset() {
	*x = Uint32{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Uint32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uint32) ProtoMessage() {}

func (x *Uint32) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Uint32) GetV() uint32 {
	if x != nil {
		return x.xxx_hidden_V
	}
	return 0
}

func (x *Uint32) SetV(v uint32) {
	x.xxx_hidden_V = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Uint32) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Uint32) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = 0
}

type Uint32_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *uint32
}

func (b0 Uint32_builder) Build() *Uint32 {
	m0 := &Uint32{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = *b.V
	}
	return m0
}

type Uint64 struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           uint64                 `protobuf:"varint,1,req,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Uint64) Reset() {
	*x = Uint64{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Uint64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Uint64) ProtoMessage() {}

func (x *Uint64) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Uint64) GetV() uint64 {
	if x != nil {
		return x.xxx_hidden_V
	}
	return 0
}

func (x *Uint64) SetV(v uint64) {
	x.xxx_hidden_V = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Uint64) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Uint64) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = 0
}

type Uint64_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *uint64
}

func (b0 Uint64_builder) Build() *Uint64 {
	m0 := &Uint64{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = *b.V
	}
	return m0
}

type Sint32 struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           int32                  `protobuf:"zigzag32,1,req,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Sint32) Reset() {
	*x = Sint32{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sint32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sint32) ProtoMessage() {}

func (x *Sint32) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Sint32) GetV() int32 {
	if x != nil {
		return x.xxx_hidden_V
	}
	return 0
}

func (x *Sint32) SetV(v int32) {
	x.xxx_hidden_V = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Sint32) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Sint32) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = 0
}

type Sint32_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *int32
}

func (b0 Sint32_builder) Build() *Sint32 {
	m0 := &Sint32{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = *b.V
	}
	return m0
}

type Sint64 struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           int64                  `protobuf:"zigzag64,1,req,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Sint64) Reset() {
	*x = Sint64{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sint64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sint64) ProtoMessage() {}

func (x *Sint64) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Sint64) GetV() int64 {
	if x != nil {
		return x.xxx_hidden_V
	}
	return 0
}

func (x *Sint64) SetV(v int64) {
	x.xxx_hidden_V = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Sint64) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Sint64) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = 0
}

type Sint64_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *int64
}

func (b0 Sint64_builder) Build() *Sint64 {
	m0 := &Sint64{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = *b.V
	}
	return m0
}

type Fixed32 struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           uint32                 `protobuf:"fixed32,1,req,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Fixed32) Reset() {
	*x = Fixed32{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Fixed32) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fixed32) ProtoMessage() {}

func (x *Fixed32) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Fixed32) GetV() uint32 {
	if x != nil {
		return x.xxx_hidden_V
	}
	return 0
}

func (x *Fixed32) SetV(v uint32) {
	x.xxx_hidden_V = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Fixed32) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Fixed32) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = 0
}

type Fixed32_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *uint32
}

func (b0 Fixed32_builder) Build() *Fixed32 {
	m0 := &Fixed32{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = *b.V
	}
	return m0
}

type Fixed64 struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           uint64                 `protobuf:"fixed64,1,req,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Fixed64) Reset() {
	*x = Fixed64{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Fixed64) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Fixed64) ProtoMessage() {}

func (x *Fixed64) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Fixed64) GetV() uint64 {
	if x != nil {
		return x.xxx_hidden_V
	}
	return 0
}

func (x *Fixed64) SetV(v uint64) {
	x.xxx_hidden_V = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Fixed64) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Fixed64) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = 0
}

type Fixed64_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *uint64
}

func (b0 Fixed64_builder) Build() *Fixed64 {
	m0 := &Fixed64{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = *b.V
	}
	return m0
}

type Float struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           float32                `protobuf:"fixed32,1,req,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Float) Reset() {
	*x = Float{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Float) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Float) ProtoMessage() {}

func (x *Float) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Float) GetV() float32 {
	if x != nil {
		return x.xxx_hidden_V
	}
	return 0
}

func (x *Float) SetV(v float32) {
	x.xxx_hidden_V = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Float) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Float) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = 0
}

type Float_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *float32
}

func (b0 Float_builder) Build() *Float {
	m0 := &Float{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = *b.V
	}
	return m0
}

type Double struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           float64                `protobuf:"fixed64,1,req,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Double) Reset() {
	*x = Double{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Double) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Double) ProtoMessage() {}

func (x *Double) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Double) GetV() float64 {
	if x != nil {
		return x.xxx_hidden_V
	}
	return 0
}

func (x *Double) SetV(v float64) {
	x.xxx_hidden_V = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Double) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Double) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = 0
}

type Double_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *float64
}

func (b0 Double_builder) Build() *Double {
	m0 := &Double{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = *b.V
	}
	return m0
}

type Bool struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           bool                   `protobuf:"varint,1,req,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Bool) Reset() {
	*x = Bool{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bool) ProtoMessage() {}

func (x *Bool) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Bool) GetV() bool {
	if x != nil {
		return x.xxx_hidden_V
	}
	return false
}

func (x *Bool) SetV(v bool) {
	x.xxx_hidden_V = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Bool) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Bool) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = false
}

type Bool_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *bool
}

func (b0 Bool_builder) Build() *Bool {
	m0 := &Bool{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = *b.V
	}
	return m0
}

type String struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           *string                `protobuf:"bytes,1,req,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *String) Reset() {
	*x = String{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *String) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*String) ProtoMessage() {}

func (x *String) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *String) GetV() string {
	if x != nil {
		if x.xxx_hidden_V != nil {
			return *x.xxx_hidden_V
		}
		return ""
	}
	return ""
}

func (x *String) SetV(v string) {
	x.xxx_hidden_V = &v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *String) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *String) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = nil
}

type String_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *string
}

func (b0 String_builder) Build() *String {
	m0 := &String{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = b.V
	}
	return m0
}

type Bytes struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           []byte                 `protobuf:"bytes,1,req,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Bytes) Reset() {
	*x = Bytes{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Bytes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bytes) ProtoMessage() {}

func (x *Bytes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Bytes) GetV() []byte {
	if x != nil {
		return x.xxx_hidden_V
	}
	return nil
}

func (x *Bytes) SetV(v []byte) {
	if v == nil {
		v = []byte{}
	}
	x.xxx_hidden_V = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Bytes) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Bytes) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = nil
}

type Bytes_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V []byte
}

func (b0 Bytes_builder) Build() *Bytes {
	m0 := &Bytes{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = b.V
	}
	return m0
}

type Message struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V  *Message_M             `protobuf:"bytes,1,req,name=v" json:"v,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Message) GetV() *Message_M {
	if x != nil {
		return x.xxx_hidden_V
	}
	return nil
}

func (x *Message) SetV(v *Message_M) {
	x.xxx_hidden_V = v
}

func (x *Message) HasV() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_V != nil
}

func (x *Message) ClearV() {
	x.xxx_hidden_V = nil
}

type Message_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *Message_M
}

func (b0 Message_builder) Build() *Message {
	m0 := &Message{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_V = b.V
	return m0
}

type Group struct {
	state            protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_Group *Group_Group           `protobuf:"group,1,req,name=Group,json=group" json:"group,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Group) Reset() {
	*x = Group{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group) ProtoMessage() {}

func (x *Group) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Group) GetGroup() *Group_Group {
	if x != nil {
		return x.xxx_hidden_Group
	}
	return nil
}

func (x *Group) SetGroup(v *Group_Group) {
	x.xxx_hidden_Group = v
}

func (x *Group) HasGroup() bool {
	if x == nil {
		return false
	}
	return x.xxx_hidden_Group != nil
}

func (x *Group) ClearGroup() {
	x.xxx_hidden_Group = nil
}

type Group_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	Group *Group_Group
}

func (b0 Group_builder) Build() *Group {
	m0 := &Group{}
	b, x := &b0, m0
	_, _ = b, x
	x.xxx_hidden_Group = b.Group
	return m0
}

type Message_M struct {
	state         protoimpl.MessageState `protogen:"opaque.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message_M) Reset() {
	*x = Message_M{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message_M) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_M) ProtoMessage() {}

func (x *Message_M) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

type Message_M_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

}

func (b0 Message_M_builder) Build() *Message_M {
	m0 := &Message_M{}
	b, x := &b0, m0
	_, _ = b, x
	return m0
}

type Group_Group struct {
	state                  protoimpl.MessageState `protogen:"opaque.v1"`
	xxx_hidden_V           int32                  `protobuf:"varint,1,opt,name=v" json:"v,omitempty"`
	XXX_raceDetectHookData protoimpl.RaceDetectHookData
	XXX_presence           [1]uint32
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *Group_Group) Reset() {
	*x = Group_Group{}
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[16]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Group_Group) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Group_Group) ProtoMessage() {}

func (x *Group_Group) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes[16]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Group_Group) GetV() int32 {
	if x != nil {
		return x.xxx_hidden_V
	}
	return 0
}

func (x *Group_Group) SetV(v int32) {
	x.xxx_hidden_V = v
	protoimpl.X.SetPresent(&(x.XXX_presence[0]), 0, 1)
}

func (x *Group_Group) HasV() bool {
	if x == nil {
		return false
	}
	return protoimpl.X.Present(&(x.XXX_presence[0]), 0)
}

func (x *Group_Group) ClearV() {
	protoimpl.X.ClearPresent(&(x.XXX_presence[0]), 0)
	x.xxx_hidden_V = 0
}

type Group_Group_builder struct {
	_ [0]func() // Prevents comparability and use of unkeyed literals for the builder.

	V *int32
}

func (b0 Group_Group_builder) Build() *Group_Group {
	m0 := &Group_Group{}
	b, x := &b0, m0
	_, _ = b, x
	if b.V != nil {
		protoimpl.X.SetPresentNonAtomic(&(x.XXX_presence[0]), 0, 1)
		x.xxx_hidden_V = *b.V
	}
	return m0
}

var File_internal_testprotos_required_required_opaque_required_opaque_proto protoreflect.FileDescriptor

var file_internal_testprotos_required_required_opaque_required_opaque_proto_rawDesc = []byte{
	0x0a, 0x42, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x2f, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2e, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x2e, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x1a, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x5f, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x05, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x12, 0x13, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x05,
	0xaa, 0x01, 0x02, 0x08, 0x03, 0x52, 0x01, 0x76, 0x22, 0x1c, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x12, 0x13, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x05, 0xaa, 0x01,
	0x02, 0x08, 0x03, 0x52, 0x01, 0x76, 0x22, 0x1d, 0x0a, 0x06, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x12, 0x13, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x05, 0xaa, 0x01, 0x02,
	0x08, 0x03, 0x52, 0x01, 0x76, 0x22, 0x1d, 0x0a, 0x06, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12,
	0x13, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08,
	0x03, 0x52, 0x01, 0x76, 0x22, 0x1d, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x13,
	0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x11, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x03,
	0x52, 0x01, 0x76, 0x22, 0x1d, 0x0a, 0x06, 0x53, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x13, 0x0a,
	0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x12, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x03, 0x52,
	0x01, 0x76, 0x22, 0x1e, 0x0a, 0x07, 0x46, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x12, 0x13, 0x0a,
	0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x07, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x03, 0x52,
	0x01, 0x76, 0x22, 0x1e, 0x0a, 0x07, 0x46, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x12, 0x13, 0x0a,
	0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x03, 0x52,
	0x01, 0x76, 0x22, 0x1c, 0x0a, 0x05, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x13, 0x0a, 0x01, 0x76,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x03, 0x52, 0x01, 0x76,
	0x22, 0x1d, 0x0a, 0x06, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x12, 0x13, 0x0a, 0x01, 0x76, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x03, 0x52, 0x01, 0x76, 0x22,
	0x1b, 0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x13, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x03, 0x52, 0x01, 0x76, 0x22, 0x1d, 0x0a, 0x06,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x13, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x05, 0xaa, 0x01, 0x02, 0x08, 0x03, 0x52, 0x01, 0x76, 0x22, 0x1c, 0x0a, 0x05, 0x42,
	0x79, 0x74, 0x65, 0x73, 0x12, 0x13, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x05, 0xaa, 0x01, 0x02, 0x08, 0x03, 0x52, 0x01, 0x76, 0x22, 0x51, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x41, 0x0a, 0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x42, 0x05, 0xaa,
	0x01, 0x02, 0x08, 0x03, 0x52, 0x01, 0x76, 0x1a, 0x03, 0x0a, 0x01, 0x4d, 0x22, 0x6d, 0x0a, 0x05,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x4d, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x2e, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x42, 0x07, 0xaa, 0x01, 0x04, 0x08, 0x03, 0x28, 0x02, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x1a, 0x15, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0c, 0x0a,
	0x01, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x76, 0x42, 0x51, 0x5a, 0x47, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x5f,
	0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x92, 0x03, 0x05, 0xd2, 0x3e, 0x02, 0x10, 0x03, 0x62, 0x08,
	0x65, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x70, 0xe8, 0x07,
}

var file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes = make([]protoimpl.MessageInfo, 17)
var file_internal_testprotos_required_required_opaque_required_opaque_proto_goTypes = []any{
	(*Int32)(nil),       // 0: opaque.goproto.proto.testrequired.Int32
	(*Int64)(nil),       // 1: opaque.goproto.proto.testrequired.Int64
	(*Uint32)(nil),      // 2: opaque.goproto.proto.testrequired.Uint32
	(*Uint64)(nil),      // 3: opaque.goproto.proto.testrequired.Uint64
	(*Sint32)(nil),      // 4: opaque.goproto.proto.testrequired.Sint32
	(*Sint64)(nil),      // 5: opaque.goproto.proto.testrequired.Sint64
	(*Fixed32)(nil),     // 6: opaque.goproto.proto.testrequired.Fixed32
	(*Fixed64)(nil),     // 7: opaque.goproto.proto.testrequired.Fixed64
	(*Float)(nil),       // 8: opaque.goproto.proto.testrequired.Float
	(*Double)(nil),      // 9: opaque.goproto.proto.testrequired.Double
	(*Bool)(nil),        // 10: opaque.goproto.proto.testrequired.Bool
	(*String)(nil),      // 11: opaque.goproto.proto.testrequired.String
	(*Bytes)(nil),       // 12: opaque.goproto.proto.testrequired.Bytes
	(*Message)(nil),     // 13: opaque.goproto.proto.testrequired.Message
	(*Group)(nil),       // 14: opaque.goproto.proto.testrequired.Group
	(*Message_M)(nil),   // 15: opaque.goproto.proto.testrequired.Message.M
	(*Group_Group)(nil), // 16: opaque.goproto.proto.testrequired.Group.Group
}
var file_internal_testprotos_required_required_opaque_required_opaque_proto_depIdxs = []int32{
	15, // 0: opaque.goproto.proto.testrequired.Message.v:type_name -> opaque.goproto.proto.testrequired.Message.M
	16, // 1: opaque.goproto.proto.testrequired.Group.group:type_name -> opaque.goproto.proto.testrequired.Group.Group
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_internal_testprotos_required_required_opaque_required_opaque_proto_init() }
func file_internal_testprotos_required_required_opaque_required_opaque_proto_init() {
	if File_internal_testprotos_required_required_opaque_required_opaque_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_testprotos_required_required_opaque_required_opaque_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   17,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_required_required_opaque_required_opaque_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_required_required_opaque_required_opaque_proto_depIdxs,
		MessageInfos:      file_internal_testprotos_required_required_opaque_required_opaque_proto_msgTypes,
	}.Build()
	File_internal_testprotos_required_required_opaque_required_opaque_proto = out.File
	file_internal_testprotos_required_required_opaque_required_opaque_proto_rawDesc = nil
	file_internal_testprotos_required_required_opaque_required_opaque_proto_goTypes = nil
	file_internal_testprotos_required_required_opaque_required_opaque_proto_depIdxs = nil
}

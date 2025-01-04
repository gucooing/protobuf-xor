// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by generate-types. DO NOT EDIT.

package impl

import (
	"reflect"

	"github.com/gucooing/protobuf-xor/reflect/protoreflect"
)

func getterForNullableScalar(fd protoreflect.FieldDescriptor, fs reflect.StructField, conv Converter, fieldOffset offset) func(p pointer) protoreflect.Value {
	ft := fs.Type
	if ft.Kind() == reflect.Ptr {
		ft = ft.Elem()
	}
	if fd.Kind() == protoreflect.EnumKind {
		elemType := fs.Type.Elem()
		// Enums for nullable types.
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			rv := p.Apply(fieldOffset).Elem().AsValueOf(elemType)
			if rv.IsNil() {
				return conv.Zero()
			}
			return conv.PBValueOf(rv.Elem())
		}
	}
	switch ft.Kind() {
	case reflect.Bool:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).BoolPtr()
			if *x == nil {
				return conv.Zero()
			}
			return protoreflect.ValueOfBool(**x)
		}
	case reflect.Int32:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Int32Ptr()
			if *x == nil {
				return conv.Zero()
			}
			return protoreflect.ValueOfInt32(**x)
		}
	case reflect.Uint32:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Uint32Ptr()
			if *x == nil {
				return conv.Zero()
			}
			return protoreflect.ValueOfUint32(**x)
		}
	case reflect.Int64:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Int64Ptr()
			if *x == nil {
				return conv.Zero()
			}
			return protoreflect.ValueOfInt64(**x)
		}
	case reflect.Uint64:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Uint64Ptr()
			if *x == nil {
				return conv.Zero()
			}
			return protoreflect.ValueOfUint64(**x)
		}
	case reflect.Float32:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Float32Ptr()
			if *x == nil {
				return conv.Zero()
			}
			return protoreflect.ValueOfFloat32(**x)
		}
	case reflect.Float64:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Float64Ptr()
			if *x == nil {
				return conv.Zero()
			}
			return protoreflect.ValueOfFloat64(**x)
		}
	case reflect.String:
		if fd.Kind() == protoreflect.BytesKind {
			return func(p pointer) protoreflect.Value {
				if p.IsNil() {
					return conv.Zero()
				}
				x := p.Apply(fieldOffset).StringPtr()
				if *x == nil {
					return conv.Zero()
				}
				if len(**x) == 0 {
					return protoreflect.ValueOfBytes(nil)
				}
				return protoreflect.ValueOfBytes([]byte(**x))
			}
		}
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).StringPtr()
			if *x == nil {
				return conv.Zero()
			}
			return protoreflect.ValueOfString(**x)
		}
	case reflect.Slice:
		if fd.Kind() == protoreflect.StringKind {
			return func(p pointer) protoreflect.Value {
				if p.IsNil() {
					return conv.Zero()
				}
				x := p.Apply(fieldOffset).Bytes()
				if len(*x) == 0 {
					return conv.Zero()
				}
				return protoreflect.ValueOfString(string(*x))
			}
		}
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Bytes()
			if *x == nil {
				return conv.Zero()
			}
			return protoreflect.ValueOfBytes(*x)
		}
	}
	panic("unexpected protobuf kind: " + ft.Kind().String())
}

func getterForDirectScalar(fd protoreflect.FieldDescriptor, fs reflect.StructField, conv Converter, fieldOffset offset) func(p pointer) protoreflect.Value {
	ft := fs.Type
	if fd.Kind() == protoreflect.EnumKind {
		// Enums for non nullable types.
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			rv := p.Apply(fieldOffset).AsValueOf(fs.Type).Elem()
			return conv.PBValueOf(rv)
		}
	}
	switch ft.Kind() {
	case reflect.Bool:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Bool()
			return protoreflect.ValueOfBool(*x)
		}
	case reflect.Int32:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Int32()
			return protoreflect.ValueOfInt32(*x)
		}
	case reflect.Uint32:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Uint32()
			return protoreflect.ValueOfUint32(*x)
		}
	case reflect.Int64:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Int64()
			return protoreflect.ValueOfInt64(*x)
		}
	case reflect.Uint64:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Uint64()
			return protoreflect.ValueOfUint64(*x)
		}
	case reflect.Float32:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Float32()
			return protoreflect.ValueOfFloat32(*x)
		}
	case reflect.Float64:
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Float64()
			return protoreflect.ValueOfFloat64(*x)
		}
	case reflect.String:
		if fd.Kind() == protoreflect.BytesKind {
			return func(p pointer) protoreflect.Value {
				if p.IsNil() {
					return conv.Zero()
				}
				x := p.Apply(fieldOffset).String()
				if len(*x) == 0 {
					return protoreflect.ValueOfBytes(nil)
				}
				return protoreflect.ValueOfBytes([]byte(*x))
			}
		}
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).String()
			return protoreflect.ValueOfString(*x)
		}
	case reflect.Slice:
		if fd.Kind() == protoreflect.StringKind {
			return func(p pointer) protoreflect.Value {
				if p.IsNil() {
					return conv.Zero()
				}
				x := p.Apply(fieldOffset).Bytes()
				return protoreflect.ValueOfString(string(*x))
			}
		}
		return func(p pointer) protoreflect.Value {
			if p.IsNil() {
				return conv.Zero()
			}
			x := p.Apply(fieldOffset).Bytes()
			return protoreflect.ValueOfBytes(*x)
		}
	}
	panic("unexpected protobuf kind: " + ft.Kind().String())
}

// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package protogen

import (
	"github.com/gucooing/protobuf-xor/internal/strs"
	"github.com/gucooing/protobuf-xor/reflect/protoreflect"
)

func opaqueNewFieldHook(desc protoreflect.FieldDescriptor, field *Field) {
	field.camelCase = strs.GoCamelCase(string(desc.Name()))
}

func opaqueNewOneofHook(desc protoreflect.OneofDescriptor, oneof *Oneof) {
	oneof.camelCase = strs.GoCamelCase(string(desc.Name()))
}

func opaqueNewMessageHook(message *Message) {
	// New name mangling scheme: Add a '_' between method base
	// name (Get, Set, Clear etc) and original field name if
	// needed.  As a special case, there is one globally reserved
	// name, e.g. "Build" thet still results in actual renaming of
	// the builder field like in the old scheme.  We begin by
	// taking care of this special case.
	for _, field := range message.Fields {
		if field.camelCase == "Build" {
			field.camelCase += "_"
		}
	}

	// Then find all names of the original field names, we do not want the old scheme to affect
	// how we name things.

	camelCases := map[string]bool{}
	for _, field := range message.Fields {
		if field.Oneof != nil {
			// We add the name of the union here (potentially many times).
			camelCases[field.Oneof.camelCase] = true
		}
		// The member fields of the oneof are considered fields in the struct although
		// they are not technically there. This is to allow changing a proto2 optional
		// to a oneof with source code compatibility.
		camelCases[field.camelCase] = true

	}
	// For each field, check if any of it's methods would clash with an original field name
	for _, field := range message.Fields {
		// Every field (except the union fields, that are taken care of separately) has
		// a Get and a Set method.
		methods := []string{"Set", "Get"}
		// For explicit presence fields, we also have Has and Clear.
		if field.Desc.HasPresence() {
			methods = append(methods, "Has", "Clear")
		}
		for _, method := range methods {
			// If any method name clashes with a field name, all methods get a
			// "_" inserted between the operation and the field name.
			if camelCases[method+field.camelCase] {
				field.hasConflictHybrid = true
			}
		}
	}
	// The union names for oneofs need only have a methods prefix if there is a clash with Has, Clear or Which in
	// hybrid and opaque-v0.
	for _, field := range message.Fields {
		if field.Oneof == nil {
			continue
		}
		for _, method := range []string{"Has", "Clear", "Which"} {
			// Same logic as for regular fields - all methods get the "_" if one needs it.
			if camelCases[method+field.Oneof.camelCase] {
				field.Oneof.hasConflictHybrid = true
			}
		}
	}

}

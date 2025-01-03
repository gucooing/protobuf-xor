// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prototext_test

import (
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/gucooing/zzz/protobuf/encoding/prototext"
	"github.com/gucooing/zzz/protobuf/proto"
	"github.com/gucooing/zzz/protobuf/reflect/protoreflect"
	"github.com/gucooing/zzz/protobuf/testing/protocmp"

	testfuzzpb "github.com/gucooing/zzz/protobuf/internal/testprotos/editionsfuzztest"
)

// roundTripAndCompareProto tests if a prototext.Marshal/Unmarshal roundtrip
// preserves the contents of the message. Note: wireBytes are a protocol
// buffer wire format message, not the proto in text format. We do this because
// a random stream of bytes (e.g. generated by the fuzz engine) is more likely
// to be valid proto wire format than that it is valid text format.
func roundTripAndCompareProto(t *testing.T, wireBytes []byte, messages ...proto.Message) {
	for _, msg := range messages {
		src := msg.ProtoReflect().Type().New().Interface()

		if err := proto.Unmarshal(wireBytes, src); err != nil {
			// Ignoring invalid wire format since we want to test the prototext
			// implementation, not the wireformat implementation.
			return
		}

		// Unknown fields are not marshaled by prototext so we strip them.
		src.ProtoReflect().SetUnknown(nil)
		var ranger func(protoreflect.FieldDescriptor, protoreflect.Value) bool
		stripUnknown := func(m protoreflect.Message) {
			m.SetUnknown(nil)
			m.Range(ranger)
		}
		ranger = func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			switch {
			case fd.IsMap():
				if fd.MapValue().Message() != nil {
					v.Map().Range(func(_ protoreflect.MapKey, v protoreflect.Value) bool {
						stripUnknown(v.Message())
						return true
					})
				}
			case fd.Message() != nil:
				if fd.Cardinality() == protoreflect.Repeated {
					l := v.List()
					for i := 0; i < l.Len(); i++ {
						stripUnknown(l.Get(i).Message())
					}
				} else {
					stripUnknown(v.Message())
				}
			}
			return true
		}
		stripUnknown(src.ProtoReflect())

		textFormat, err := prototext.Marshal(src)
		if err != nil {
			t.Errorf("failed to marshal messsage to text format: %v\nmessage: %v", err, src)
		}
		dst := msg.ProtoReflect().Type().New().Interface()

		if err := (prototext.Unmarshal(textFormat, dst)); err != nil {
			t.Errorf("failed to unmarshal messsage from text format: %v\ntext format: %s", err, textFormat)
		}

		// The cmp package does not deal with NaN on its own and will report
		// NaN != NaN.
		optNaN64 := cmp.Comparer(func(x, y float32) bool {
			return (math.IsNaN(float64(x)) && math.IsNaN(float64(y))) || x == y
		})
		optNaN32 := cmp.Comparer(func(x, y float64) bool {
			return (math.IsNaN(x) && math.IsNaN(y)) || x == y
		})
		if diff := cmp.Diff(src, dst, protocmp.Transform(), optNaN64, optNaN32); diff != "" {
			t.Error(diff)
		}
	}
}

func FuzzEncodeDecodeRoundTrip(f *testing.F) {
	f.Add([]byte("Hello World!"))
	f.Fuzz(func(t *testing.T, in []byte) {
		roundTripAndCompareProto(t, in, (*testfuzzpb.TestAllTypesProto2)(nil), (*testfuzzpb.TestAllTypesProto2Editions)(nil), (*testfuzzpb.TestAllTypesProto3)(nil), (*testfuzzpb.TestAllTypesProto3Editions)(nil))
	})
}

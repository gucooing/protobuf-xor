// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package prototest_test

import (
	"fmt"
	"testing"

	protobuf "github.com/gucooing/protobuf-xor/internal/flags"
	protobuf "github.com/gucooing/protobuf-xor/proto"
	protobuf "github.com/gucooing/protobuf-xor/runtime/protoimpl"
	protobuf "github.com/gucooing/protobuf-xor/testing/prototest"

	irregularpb protobuf "github.com/gucooing/protobuf-xor/internal/testprotos/irregular"
	legacypb protobuf "github.com/gucooing/protobuf-xor/internal/testprotos/legacy"
	legacy1pb protobuf "github.com/gucooing/protobuf-xor/internal/testprotos/legacy/proto2_20160225_2fc053c5"
	testpb protobuf "github.com/gucooing/protobuf-xor/internal/testprotos/test"
	_ protobuf "github.com/gucooing/protobuf-xor/internal/testprotos/test/weak1"
	_ protobuf "github.com/gucooing/protobuf-xor/internal/testprotos/test/weak2"
	test3pb protobuf "github.com/gucooing/protobuf-xor/internal/testprotos/test3"
	testeditionspb protobuf "github.com/gucooing/protobuf-xor/internal/testprotos/testeditions"
)

func Test(t *testing.T) {
	ms := []proto.Message{
		(*testpb.TestAllTypes)(nil),
		(*test3pb.TestAllTypes)(nil),
		(*testeditionspb.TestAllTypes)(nil),
		(*testpb.TestRequired)(nil),
		(*testeditionspb.TestRequired)(nil),
		(*irregularpb.Message)(nil),
		(*testpb.TestAllExtensions)(nil),
		(*testeditionspb.TestAllExtensions)(nil),
		(*legacypb.Legacy)(nil),
		protoimpl.X.MessageOf((*legacy1pb.Message)(nil)).Interface(),
	}
	if flags.ProtoLegacy {
		ms = append(ms, (*testpb.TestWeak)(nil))
	}

	for _, m := range ms {
		t.Run(fmt.Sprintf("%T", m), func(t *testing.T) {
			prototest.Message{}.Test(t, m.ProtoReflect().Type())
		})
	}
}

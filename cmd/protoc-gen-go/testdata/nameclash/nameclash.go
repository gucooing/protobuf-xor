// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package nameclash is imported by gen_test.go and itself imports the various
// nameclash proto variants.
package nameclash

import (
	_ protobuf "github.com/gucooing/protobuf-xor/cmd/protoc-gen-go/testdata/nameclash/test_name_clash_hybrid"
	_ protobuf "github.com/gucooing/protobuf-xor/cmd/protoc-gen-go/testdata/nameclash/test_name_clash_hybrid3"
	_ protobuf "github.com/gucooing/protobuf-xor/cmd/protoc-gen-go/testdata/nameclash/test_name_clash_opaque"
	_ protobuf "github.com/gucooing/protobuf-xor/cmd/protoc-gen-go/testdata/nameclash/test_name_clash_opaque3"
	_ protobuf "github.com/gucooing/protobuf-xor/cmd/protoc-gen-go/testdata/nameclash/test_name_clash_open"
	_ protobuf "github.com/gucooing/protobuf-xor/cmd/protoc-gen-go/testdata/nameclash/test_name_clash_open3"
)

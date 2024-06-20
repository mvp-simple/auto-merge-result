// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The protoreflect build tag disables use of fast-path methods.
//go:build !protoreflect
// +build !protoreflect

package proto

import (
	"belik.org/protobuf/reflect/protoreflect"
	"belik.org/protobuf/runtime/protoiface"
)

const hasProtoMethods = true

func protoMethods(m protoreflect.Message) *protoiface.Methods {
	return m.ProtoMethods()
}

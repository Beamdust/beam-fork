// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package funcx

import (
	"reflect"
	"testing"

	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/core/typex"
	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/core/util/reflectx"
)

func TestSatisfy(t *testing.T) {
	tests := []struct {
		Sig *Signature
		Fn  any
		Ok  bool
	}{
		// Concrete signature
		{
			Sig: MakePredicate(reflectx.Int),
			Fn:  func(int) bool { return true },
			Ok:  true,
		},
		{
			Sig: MakePredicate(reflectx.Int),
			Fn:  func(string) bool { return true },
			Ok:  false, // wrong type
		},
		{
			Sig: MakePredicate(reflectx.Int),
			Fn:  func(typex.T) bool { return true },
			Ok:  true,
		},
		{
			Sig: MakePredicate(reflectx.Int),
			Fn:  func(typex.X) bool { return true },
			Ok:  true,
		},
		{
			Sig: MakePredicate(reflectx.Int),
			Fn:  func(typex.X) (bool, error) { return true, nil },
			Ok:  false, // extra return
		},
		{
			Sig: MakePredicate(reflectx.Int),
			Fn:  func(int, typex.X) bool { return true },
			Ok:  false, // extra parameter
		},
		{
			Sig: MakePredicate(reflectx.Int, reflectx.String),
			Fn:  func(int, string) bool { return true },
			Ok:  true,
		},
		{
			Sig: MakePredicate(reflectx.Int, reflectx.String),
			Fn:  func(int, typex.T) bool { return true },
			Ok:  true,
		},
		{
			Sig: MakePredicate(reflectx.Int, reflectx.String),
			Fn:  func(typex.T, typex.T) bool { return true },
			Ok:  false, // bind conflict
		},
		// Concrete signatures with optionals
		{
			Sig: &Signature{
				OptArgs: []reflect.Type{reflectx.Int},
				Args:    []reflect.Type{reflectx.String, reflectx.Int},
			},
			Fn: func(string, int) {},
			Ok: true,
		},
		{
			Sig: &Signature{
				OptArgs: []reflect.Type{reflectx.Int},
				Args:    []reflect.Type{reflectx.String, reflectx.Int},
			},
			Fn: func(int, string, int) {},
			Ok: true,
		},
		{
			Sig: &Signature{
				OptArgs: []reflect.Type{reflectx.Int},
				Args:    []reflect.Type{reflectx.String, reflectx.Int},
			},
			Fn: func(string) {},
			Ok: false, // too few parameters
		},
		{
			Sig: &Signature{
				OptArgs: []reflect.Type{reflectx.Int},
				Args:    []reflect.Type{reflectx.String, reflectx.Int},
			},
			Fn: func(int, string) {},
			Ok: false, // type mismatch
		},
		{
			Sig: &Signature{
				OptArgs: []reflect.Type{reflectx.Int},
				Args:    []reflect.Type{reflectx.String, reflectx.Int},
			},
			Fn: func(int, int, string, int) {},
			Ok: false, // too many parameters
		},
		{
			Sig: &Signature{
				OptArgs: []reflect.Type{reflectx.Int, reflectx.String},
				Args:    []reflect.Type{reflectx.String, reflectx.Int},
			},
			Fn: func(typex.T, string, typex.T) {},
			Ok: true, // Subtle: matches the first opt param
		},
		{
			Sig: &Signature{
				OptArgs: []reflect.Type{reflectx.Int, reflectx.String},
				Args:    []reflect.Type{reflectx.String, reflectx.Int},
			},
			Fn: func(typex.T, typex.T, int) {},
			Ok: true, // Subtle: matches the second opt param
		},
		{
			Sig: &Signature{
				OptArgs: []reflect.Type{reflectx.Int, reflectx.String},
				Args:    []reflect.Type{reflectx.String, reflectx.Int},
			},
			Fn: func(typex.T, typex.X, int) {},
			Ok: false, // T is unbound
		},
		{
			Sig: &Signature{
				Return:    []reflect.Type{reflectx.Int, reflectx.String},
				OptReturn: []reflect.Type{reflectx.Int, reflectx.String},
			},
			Fn: func() (int, string, int) { return 0, "", 0 },
			Ok: true, // Subtle: matches the first opt return
		},
		{
			Sig: &Signature{
				Return:    []reflect.Type{reflectx.Int, reflectx.String},
				OptReturn: []reflect.Type{reflectx.Int, reflectx.String},
			},
			Fn: func() (int, string, string) { return 0, "", "" },
			Ok: true, // Subtle: matches the second opt return
		},
		{
			Sig: &Signature{
				OptArgs:   []reflect.Type{reflectx.Int, reflectx.String},
				Args:      []reflect.Type{reflectx.String, reflectx.Int},
				Return:    []reflect.Type{reflectx.Int, reflectx.String},
				OptReturn: []reflect.Type{reflectx.Int, reflectx.String},
			},
			Fn: func(int, typex.X, typex.Y) (typex.Y, string, typex.X) { return nil, "", nil },
			Ok: true, // Subtle: X -> string, Y -> int, so matches second opt return
		},
		// Generic signatures
		{
			Sig: MakePredicate(typex.TType),
			Fn:  func(typex.T) bool { return true },
			Ok:  true,
		},
		{
			Sig: MakePredicate(typex.TType),
			Fn:  func(typex.X) bool { return true },
			Ok:  true,
		},
		{
			Sig: MakePredicate(typex.TType, typex.UType, reflectx.String),
			Fn:  func(typex.X, typex.Y, string) bool { return true },
			Ok:  true,
		},
		{
			Sig: MakePredicate(typex.TType, typex.UType, reflectx.String),
			Fn:  func(typex.X, typex.Y, typex.T) bool { return true },
			Ok:  true,
		},
		{
			Sig: MakePredicate(typex.TType, typex.UType, reflectx.String),
			Fn:  func(typex.X, typex.X, string) bool { return true },
			Ok:  false, // bind conflict
		},
		{
			Sig: MakePredicate(typex.TType, typex.UType, reflectx.String),
			Fn:  func(typex.X, typex.Y, typex.Y) bool { return true },
			Ok:  false, // bind conflict
		},
		{
			Sig: MakePredicate(typex.TType, typex.UType, reflectx.String),
			Fn:  func(typex.X, int, string) bool { return true },
			Ok:  false, // type mismatch
		},
	}

	for _, test := range tests {
		if err := Satisfy(test.Fn, test.Sig); (err == nil) != test.Ok {
			t.Errorf("iface: Satisfy(%v, %v) = %v, want (err==nil)==%v", reflect.ValueOf(test.Fn).Type(), test.Sig, err, test.Ok)
		}

		rfn := reflectx.MakeFunc(test.Fn)
		if err := Satisfy(rfn, test.Sig); (err == nil) != test.Ok {
			t.Errorf("reflectx.Func: Satisfy(%v, %v) = %v, want (err==nil)==%v", reflect.ValueOf(test.Fn).Type(), test.Sig, err, test.Ok)
		}

		fx, err := New(rfn)
		if err != nil {
			t.Errorf("Unable to create New Fn from reflectx.Func %v: %v", rfn.Name(), err)
			continue
		}
		if err := Satisfy(fx, test.Sig); (err == nil) != test.Ok {
			t.Errorf("*funcx.Fn:Satisfy(%v, %v) = %v, want (err==nil)==%v", reflect.ValueOf(test.Fn).Type(), test.Sig, err, test.Ok)
		}
	}
}

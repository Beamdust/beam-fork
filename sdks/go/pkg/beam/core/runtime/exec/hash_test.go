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

package exec

import (
	"encoding/json"
	"fmt"
	"hash/maphash"
	"reflect"
	"strings"
	"testing"

	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/core/graph/coder"
	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/core/graph/window"
	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/core/runtime/coderx"
	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/core/typex"
	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/core/util/reflectx"
)

func BenchmarkPrimitives(b *testing.B) {
	var value FullValue
	myHash := &maphash.Hash{}
	wfn := window.NewGlobalWindows()
	we := MakeWindowEncoder(wfn.Coder())
	b.Run("int", func(b *testing.B) {
		test := any(int(42424242))
		b.Run("native", func(b *testing.B) {
			m := make(map[int]FullValue)
			for i := 0; i < b.N; i++ {
				k := test.(int)
				value = m[k]
				m[k] = value
			}
		})
		cc, err := coderx.NewVarIntZ(reflectx.Int)
		if err != nil {
			b.Fatal(err)
		}
		encoded := &customEncodedHasher{hash: myHash, coder: makeEncoder(cc.Enc.Fn), we: we}
		// Route through constructor to init cache.
		dedicated := newNumberHasher(myHash, we)
		hashbench(b, test, encoded, dedicated)
	})
	b.Run("float32", func(b *testing.B) {
		test := any(float32(42424242.242424))
		b.Run("native", func(b *testing.B) {
			m := make(map[float32]FullValue)
			for i := 0; i < b.N; i++ {
				k := test.(float32)
				value = m[k]
				m[k] = value
			}
		})
		cc, err := coderx.NewFloat(reflectx.Float32)
		if err != nil {
			b.Fatal(err)
		}
		encoded := &customEncodedHasher{hash: myHash, coder: makeEncoder(cc.Enc.Fn), we: we}
		// Route through constructor to init cache.
		dedicated := newNumberHasher(myHash, we)
		hashbench(b, test, encoded, dedicated)
	})

	b.Run("string", func(b *testing.B) {
		tests := []any{
			"I am the very model of a modern major string.",
			"this is 10",
			strings.Repeat("100 chars!", 10),   // 100
			strings.Repeat("1k chars!!", 100),  // 1000
			strings.Repeat("10k chars!", 1000), // 10000
		}
		for _, test := range tests {
			b.Run(fmt.Sprint(len(test.(string))), func(b *testing.B) {
				b.Run("native", func(b *testing.B) {
					m := make(map[string]FullValue)
					for i := 0; i < b.N; i++ {
						k := test.(string)
						value = m[k]
						m[k] = value
					}
				})

				cc, err := coderx.NewString()
				if err != nil {
					b.Fatal(err)
				}
				encoded := &customEncodedHasher{hash: myHash, coder: makeEncoder(cc.Enc.Fn), we: we}
				dedicated := &stringHasher{hash: myHash, we: we}
				hashbench(b, test, encoded, dedicated)
			})
		}
	})
	b.Run("struct", func(b *testing.B) {
		tests := []any{
			struct {
				A      int
				B      string
				Foobar [4]int
			}{A: 56, B: "stringtastic", Foobar: [4]int{4, 2, 3, 1}},
		}
		for _, test := range tests {
			typ := reflect.TypeOf(test)
			b.Run(fmt.Sprint(typ.String()), func(b *testing.B) {
				encoded := &customEncodedHasher{hash: myHash, coder: &jsonEncoder{}, we: we}
				dedicated := &rowHasher{hash: myHash, coder: MakeElementEncoder(coder.NewR(typex.New(typ))), we: we}
				hashbench(b, test, encoded, dedicated)
			})
		}
	})
}

type jsonEncoder struct{}

func (*jsonEncoder) Encode(t reflect.Type, element any) ([]byte, error) {
	return json.Marshal(element)
}

func hashbench(b *testing.B, test any, encoded, dedicated elementHasher) {
	var value FullValue
	gw := window.SingleGlobalWindow[0]
	b.Run("interface", func(b *testing.B) {
		m := make(map[any]FullValue)
		for i := 0; i < b.N; i++ {
			k := test
			value = m[k]
			m[k] = value
		}
	})
	b.Run("encodedHash", func(b *testing.B) {
		m := make(map[uint64]FullValue)
		for i := 0; i < b.N; i++ {
			k, err := encoded.Hash(test, gw)
			if err != nil {
				b.Fatal(err)
			}
			value = m[k]
			m[k] = value
		}
	})
	if dedicated == nil {
		return
	}
	b.Run("dedicatedHash", func(b *testing.B) {
		m := make(map[uint64]FullValue)
		for i := 0; i < b.N; i++ {
			k, err := dedicated.Hash(test, gw)
			if err != nil {
				b.Fatal(err)
			}
			value = m[k]
			m[k] = value
		}
	})
}

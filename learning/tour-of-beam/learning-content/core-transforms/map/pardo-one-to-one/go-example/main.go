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

// beam-playground:
//   name: pardo-one-to-one
//   description: ParDo one-to-one example.
//   multifile: false
//   context_line: 37
//   categories:
//     - Quickstart
//   complexity: MEDIUM
//   tags:
//     - hellobeam

package main

import (
	"context"

	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam"
	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/log"
	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/x/beamx"
	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/x/debug"
)

func main() {
	ctx := context.Background()

	p, s := beam.NewPipelineWithRoot()

	// List of elements
	input := beam.Create(s, 1, 2, 3, 4, 5)

	// The applyTransform() converts [input] to [output]
	output := applyTransform(s, input)

	debug.Print(s, output)

	err := beamx.Run(ctx, p)

	if err != nil {
		log.Exitf(context.Background(), "Failed to execute job: %v", err)
	}
}

// The applyTransform() using simple multiple function
func applyTransform(s beam.Scope, input beam.PCollection) beam.PCollection {
	return beam.ParDo(s, multiplyBy10Fn, input)
}

// Multiplications by 10
func multiplyBy10Fn(element int) int {
	return element * 10
}

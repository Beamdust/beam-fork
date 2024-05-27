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
//   name: flatten
//   description: Flatten example.
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

	"github.com/Beamdust/beam-fork/go/pkg/beam"
	"github.com/Beamdust/beam-fork/go/pkg/beam/log"
	"github.com/Beamdust/beam-fork/go/pkg/beam/x/beamx"
	"github.com/Beamdust/beam-fork/go/pkg/beam/x/debug"
)

func main() {
	ctx := context.Background()

	p, s := beam.NewPipelineWithRoot()

	// List of elements start with a
	aWords := beam.Create(s, "apple", "ant", "arrow")

	// List of elements start with b
	bWords := beam.Create(s, "ball", "book", "bow")

	// The applyTransform() converts [aWords] and [bWords] to [output]
	output := applyTransform(s, aWords, bWords)

	debug.Print(s, output)

	err := beamx.Run(ctx, p)

	if err != nil {
		log.Exitf(context.Background(), "Failed to execute job: %v", err)
	}
}

// The applyTransform two PCollection data types are the same combines and returns one PCollection
func applyTransform(s beam.Scope, aInputs beam.PCollection, bInputs beam.PCollection) beam.PCollection {
	return beam.Flatten(s, aInputs, bInputs)
}

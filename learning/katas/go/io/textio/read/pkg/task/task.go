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

package task

import (
	"strings"

	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam"
	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/io/textio"
)

// Read reads from filename(s) specified by a glob string and a returns a PCollection<string>.
func Read(s beam.Scope, glob string) beam.PCollection {
	return textio.Read(s, glob)
}

// ApplyTransform converts to uppercase all elements in a PCollection<string>.
func ApplyTransform(s beam.Scope, input beam.PCollection) beam.PCollection {
	return beam.ParDo(s, strings.ToUpper, input)
}

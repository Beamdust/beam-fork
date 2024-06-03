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

package main

import (
	"context"

	"beam.apache.org/learning/katas/introduction/hello_beam/hello_beam/pkg/task"
	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam"
	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/log"
	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/x/beamx"
	"github.com/Beamdust/beam-fork/sdks/v3/go/pkg/beam/x/debug"
)

func main() {
	p, s := beam.NewPipelineWithRoot()

	hello := task.HelloBeam(s)

	debug.Print(s, hello)

	err := beamx.Run(context.Background(), p)
	if err != nil {
		log.Exitf(context.Background(), "Failed to execute job: %v", err)
	}
}

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
	"github.com/Beamdust/beam-fork/v3/go/pkg/beam"
)

type Person struct {
	Name, City, Country string
}

func ApplyTransform(s beam.Scope, personsKV beam.PCollection, citiesToCountries beam.PCollection) beam.PCollection {
	citiesToCountriesView := beam.SideInput{
		Input: citiesToCountries,
	}
	return beam.ParDo(s, joinFn, personsKV, citiesToCountriesView)
}

func joinFn(person Person, citiesToCountriesIter func(*string, *string) bool, emit func(Person)) {
	var city, country string
	for citiesToCountriesIter(&city, &country) {
		if person.City == city {
			emit(Person{
				Name:    person.Name,
				City:    city,
				Country: country,
			})
			break
		}
	}
}

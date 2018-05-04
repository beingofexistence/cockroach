// Copyright 2018 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package props_test

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/sql/opt"
	"github.com/cockroachdb/cockroach/pkg/sql/opt/props"
)

func TestPhysicalProps(t *testing.T) {
	// Empty props.
	phys := &props.Physical{}
	testPhysicalProps(t, phys, "")

	if phys.Defined() {
		t.Error("no props should be defined")
	}

	// Presentation props.
	presentation := props.Presentation{
		opt.LabeledColumn{Label: "a", ID: 1},
		opt.LabeledColumn{Label: "b", ID: 2},
	}
	phys = &props.Physical{Presentation: presentation}
	testPhysicalProps(t, phys, "p:a:1,b:2")

	if !presentation.Defined() {
		t.Error("presentation should be defined")
	}

	if !presentation.Equals(presentation) {
		t.Error("presentation should equal itself")
	}

	if presentation.Equals(props.Presentation{}) {
		t.Error("presentation should not equal the empty presentation")
	}

	// Add Ordering props.
	ordering := props.Ordering{1, 5}
	phys.Ordering = ordering
	testPhysicalProps(t, phys, "p:a:1,b:2 o:+1,+5")

	if !ordering.Defined() {
		t.Error("ordering should be defined")
	}

	if !ordering.Provides(ordering) {
		t.Error("ordering should provide itself")
	}

	if !ordering.Provides(props.Ordering{1}) {
		t.Error("ordering should provide the prefix ordering")
	}

	if (props.Ordering{}).Provides(ordering) {
		t.Error("empty ordering should not provide ordering")
	}

	if !ordering.Provides(props.Ordering{}) {
		t.Error("ordering should provide the empty ordering")
	}

	if !ordering.Equals(ordering) {
		t.Error("ordering should be equal with itself")
	}

	if ordering.Equals(props.Ordering{}) {
		t.Error("ordering should not equal the empty ordering")
	}

	if (props.Ordering{}).Equals(ordering) {
		t.Error("empty ordering should not equal ordering")
	}
}

func testPhysicalProps(t *testing.T, physProps *props.Physical, expected string) {
	actual := physProps.Fingerprint()
	if actual != expected {
		t.Errorf("\nexpected: %s\nactual: %s", expected, actual)
	}
}

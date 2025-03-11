// Copyright 2015 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package relaxng

import (
	"testing"
)

func TestRemoveTODO(t *testing.T) {
	g := &Grammar{
		Start: &NameOrPattern{
			Value: &Value{
				Ns:   "TODO",
				Text: "Hello",
			},
		},
		Define: []Define{
			{
				Element: Pair{
					Right: &NameOrPattern{
						Value: &Value{
							Ns:   "TODO",
							Text: "Hello",
						},
					},
				},
			},
		},
	}
	RemoveTODOs(g)
	if g.Start.Value.Ns != "" {
		t.Fatalf("Ns not cleared")
	}
	if g.Start.Value.Text != "Hello" {
		t.Fatalf("Hello is where?")
	}
	if g.Define[0].Element.Right.Value.Ns != "" {
		t.Fatalf("Define Ns not cleared")
	}
}

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
	"fmt"
)

// Parses the simplified RelaxNG Grammar,
// translates it to relapse and then validates the input xml.
func ExampleValidate() {
	simplifiedRelaxNG := `
	<grammar>
	    <start>
	        <ref name="element1"></ref>
	    </start>
	    <define name="element1">
	        <element>
	            <name>foo</name>
	            <empty></empty>
	        </element>
	    </define>
	</grammar>`
	relaxing, _ := ParseGrammar([]byte(simplifiedRelaxNG))
	relapse, _ := Translate(relaxing)
	input := "<foo/>"
	if err := Validate(relapse, []byte(input)); err != nil {
		fmt.Println("invalid")
	}
	fmt.Println("valid")
	// Output: valid
}

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
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenJSSuite(t *testing.T) {
	tests := scanFiles()
	fmt.Fprintf(os.Stderr, "var tests = [\n")
	for _, test := range tests {
		tbase := filepath.Base(filepath.Dir(test.Filename))
		for _, c := range test.Xmls {
			cbases := strings.Split(filepath.Base(c.Filename), ".")
			cbase := strings.Join(cbases[:len(cbases)-1], ".")
			fmt.Fprintf(os.Stderr, "'%s.%s',\n", tbase, cbase)
		}
	}
	fmt.Fprintf(os.Stderr, "]\n")
}

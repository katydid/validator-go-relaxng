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
	"strings"
	"testing"
)

func TestSimpleParseSpecExample(t *testing.T) {
	//http://relaxng.org/spec-20011203.html
	example5p1 := `<?xml version="1.0"?>
<grammar xmlns="http://relaxng.org/ns/structure/1.0">
  <start>
    <ref name="foo.element"/>
  </start>

  <define name="foo.element">
    <element>
      <name ns="">foo</name>
      <group>
        <ref name="bar1.element"/>
        <ref name="bar2.element"/>
      </group>
    </element>
  </define>

  <define name="bar1.element">
    <element>
      <name ns="http://www.example.com/n1">bar1</name>
      <empty/>
    </element>
  </define>

  <define name="bar2.element">
    <element>
      <name ns="http://www.example.com/n2">bar2</name>
      <empty/>
    </element>
  </define>
</grammar>`
	g, err := ParseGrammar([]byte(example5p1))
	if err != nil {
		panic(err)
	}
	s := g.String()
	t.Logf("%s", s)
	if !strings.Contains(s, `ref name="bar2.element"`) {
		t.Fatalf("expected ref name bar2.element in group")
	}
	if !strings.Contains(s, `ref name="bar1.element"`) {
		t.Fatalf("expected ref name bar1.element in group")
	}
	if !strings.Contains(s, `<define name="bar1.element">`) {
		t.Fatalf("expected define bar1")
	}
}

func TestSimpleParseChoice(t *testing.T) {
	choice := `<grammar xmlns="http://relaxng.org/ns/structure/1.0">
  <start>
    <ref name="element1"/>
  </start>
  <define name="element1">
    <element>
    <name>foo</name>
    <choice> 
      <ref name="a"/>
      <ref name="b"/>
    </choice>
    </element>
  </define>
</grammar>`
	g, err := ParseGrammar([]byte(choice))
	if err != nil {
		t.Fatal(err)
	}
	s := g.String()
	t.Log(s)
	if !strings.Contains(s, `choice`) {
		t.Fatalf("expected choice")
	}
	if !strings.Contains(s, `ref name="a"`) {
		t.Fatalf("expected ref name a in choice")
	}
}

func TestSimpleParse94(t *testing.T) {
	s94 := `<grammar xmlns="http://relaxng.org/ns/structure/1.0">
  <start>
    <ref name="element1"/>
  </start>
  <define name="element1">
    <element>
      <name ns="">foo</name>
      <attribute>
        <name ns="">bar</name>
        <text/>
      </attribute>
    </element>
  </define>
</grammar>`
	g, err := ParseGrammar([]byte(s94))
	if err != nil {
		t.Fatal(err)
	}
	s := g.String()
	t.Log(s)
	if !strings.Contains(s, `attribute`) {
		t.Fatalf("expected attribute")
	}
}

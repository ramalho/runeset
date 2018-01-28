package runeset

import (
	"fmt"
	"reflect"
	"testing"
)

const (
	universe  = "0123456789"
	even      = "02468"
	odd       = "13579"
	prime     = "2357"
	fibonacci = "012358"
	singleton = "1"
	void      = ""
)

var (
	U = MakeFromString(universe)
	E = MakeFromString(even)
	O = MakeFromString(odd)
	P = MakeFromString(prime)
	F = MakeFromString(fibonacci)
	V = MakeFromString(void)
)

func TestHas_empty(t *testing.T) {
	var s Set
	if s.Has('x') {
		t.Errorf(`s is empty, no words should be found.`)
	}
}

func TestHas(t *testing.T) {
	if !E.Has('2') {
		t.Errorf(`2 should be in E (even)`)
	}
}

func TestAdd_empty(t *testing.T) {
	s := Set{}
	r := 'x'
	s.Add(r)
	if res := !s.Has(r); res {
		t.Errorf(`%q was added, but s.Has(%[1]q) == %v.`, r, res)
	}
}

func TestMake(t *testing.T) {
	chars := []rune("abc")
	s := Make(chars...)
	for _, c := range chars {
		if res := !s.Has(c); res {
			t.Errorf(`%q was added, but s.Has(%[1]q) == %v.`, c, res)
		}
	}

}

func TestMakeFromString(t *testing.T) {
	chars := "abc"
	s := MakeFromString(chars)
	for _, c := range chars {
		if res := !s.Has(c); res {
			t.Errorf(`%q was added, but s.Has(%[1]q) == %v.`, c, res)
		}
	}

}

func Example() {
	s1 := MakeFromString("ABCD")
	s2 := MakeFromString("BDEF")
	fmt.Println(s1.Intersection(s2))
	// Output: Set{B D}
}

func TestIntersection(t *testing.T) {
	s1 := MakeFromString("abcd")
	s2 := MakeFromString("bdz")
	want := MakeFromString("bd")
	got := s1.Intersection(s2)
	if !reflect.DeepEqual(want, got) {
		t.Errorf(`Wanted: %v Got: %v`, want, got)
	}

}

func TestIntersection_empty(t *testing.T) {
	s1 := MakeFromString("abcd")
	s2 := MakeFromString("xyz")
	want := Set{}
	got := s1.Intersection(s2)
	if !reflect.DeepEqual(want, got) {
		t.Errorf(`Wanted: %v Got: %v`, want, got)
	}

}

func TestIntersection_table(t *testing.T) {
	testCases := []struct {
		name   string
		first  Set
		second Set
		want   Set
	}{
		{"even & prime", E, P, MakeFromString("2")},
		{"even & fibonacci", E, F, MakeFromString("028")},
		{"even & odd", E, O, V},
		{"void & universe", V, U, V},
		{"even & universe", E, U, E},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.first.Intersection(tc.second)
			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}

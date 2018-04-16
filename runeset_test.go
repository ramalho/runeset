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
	S = MakeFromString(singleton)
	V = MakeFromString(void)
)

func TestHas_empty(t *testing.T) {
	var s Set
	if s.Contains('x') {
		t.Errorf(`s is empty, no words should be found.`)
	}
}

func TestHas(t *testing.T) {
	if !E.Contains('2') {
		t.Errorf(`2 should be in E (even)`)
	}
}

func TestAdd_empty(t *testing.T) {
	s := Set{}
	r := 'x'
	s.Add(r)
	if res := !s.Contains(r); res {
		t.Errorf(`%q was added, but s.Contains(%[1]q) == %v.`, r, res)
	}
}

func TestMake(t *testing.T) {
	chars := []rune("abc")
	s := Make(chars...)
	for _, c := range chars {
		if res := !s.Contains(c); res {
			t.Errorf(`%q was added, but s.Contains(%[1]q) == %v.`, c, res)
		}
	}

}

func TestMakeFromString(t *testing.T) {
	chars := "abc"
	s := MakeFromString(chars)
	for _, c := range chars {
		if res := !s.Contains(c); res {
			t.Errorf(`%q was added, but s.Contains(%[1]q) == %v.`, c, res)
		}
	}

}

func TestSorted(t *testing.T) {
	s := MakeFromString("BADC")
	got := s.Sorted()
	want := []rune{'A', 'B', 'C', 'D'}
	if !reflect.DeepEqual(want, got) {
		t.Errorf(`Wanted: %v Got: %v`, want, got)
	}
}

func Example() {
	s1 := MakeFromString("1234567")
	s2 := MakeFromString("86420")
	fmt.Println(s1.Intersection(s2))
	// Output: Set{2 4 6}
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
		{"sigleton & void", S, V, V},
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

func TestIntersectionUpdate(t *testing.T) {
	testCases := []struct {
		name   string
		receiver  Set
		other  Set
		want   Set
	}{
		{"even & prime", MakeFromString(even), MakeFromString(prime), MakeFromString("2")},
		{"even & fibonacci", MakeFromString(even), MakeFromString(fibonacci), MakeFromString("028")},
		{"even & odd", MakeFromString(even), MakeFromString(odd), Set{}},
		{"void & universe", Set{}, MakeFromString(universe), Set{}},
		{"even & universe", MakeFromString(even), MakeFromString(universe), MakeFromString(even)},
		{"sigleton & void", MakeFromString(singleton), Set{}, Set{}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.receiver.IntersectionUpdate(tc.other)
			if !reflect.DeepEqual(tc.want, tc.receiver) {
				t.Errorf("got %v; want %v", tc.receiver, tc.want)
			}
		})
	}
}

func TestCopy(t *testing.T) {
	testCases := []Set{
		Make(),
		Make('a'),
		Make('a', 'b'),
	}
	for _, set := range testCases {
		t.Run(fmt.Sprintf("%v.Copy()", set), func(t *testing.T) {
			clone := set.Copy()
			if !reflect.DeepEqual(set, clone) {
				t.Errorf("clone: %v; original: %v", clone, set)
			}
			set['z'] = struct{}{}
			if reflect.DeepEqual(set, clone) {
				t.Errorf("clone: %v; original: %v", clone, set)
			}
		})
	}
}

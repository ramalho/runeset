package runeset

import (
	"bytes"
	"sort"
)

// Set represents a set of runes
type Set map[rune]struct{}

// Make creates and returns a new Set
func Make(chars ...rune) Set {
	s := Set{}
	for _, c := range chars {
		s.Add(c)
	}
	return s
}

// MakeFromString creates and returns a new Set
func MakeFromString(chars string) Set {
	s := Set{}
	for _, c := range chars {
		s.Add(c)
	}
	return s
}

func (s Set) String() string {
	var codes []int
	for c := range s {
		codes = append(codes, int(c))
	}
	sort.Ints(codes)
	var buf bytes.Buffer
	buf.WriteString("Set{")
	for i, c := range codes {
		if i > 0 {
			buf.WriteByte(' ')
		}
		buf.WriteString(string(c))
	}
	buf.WriteByte('}')
	return buf.String()
}

// Add adds a rune to the set.
func (s Set) Add(r rune) {
	s[r] = struct{}{} // zero-byte struct
}

// Has reports whether set contains given rune
func (s Set) Has(r rune) bool {
	_, found := s[r]
	return found
}

// Intersection returns a new set: the intersection of s and other
func (s Set) Intersection(other Set) Set {
	result := Set{}
	if len(other) > 0 {
		for r := range s {
			if other.Has(r) {
				result.Add(r)
			}
		}
	}
	return result
}

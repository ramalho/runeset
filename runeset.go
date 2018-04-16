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

type runeSlice []rune

func (rs runeSlice) Len() int           { return len(rs) }
func (rs runeSlice) Less(i, j int) bool { return rs[i] < rs[j] }
func (rs runeSlice) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }

// Sorted returns sorted slice with runes from s
func (s Set) Sorted() []rune {
	var rs []rune
	for c := range s {
		rs = append(rs, c)
	}
	sort.Sort(runeSlice(rs))
	return rs
}

func (s Set) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	for i, c := range s.Sorted() {
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

// Intersection returns a new set: the intersection of s AND other
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

// IntersectionUpdate changes receiver in-place, keeping only
// elements that are in it AND in other.
func (s Set) IntersectionUpdate(other Set) {
	for r := range s {
		if !other.Has(r) {
			delete(s, r)
		}
	}
}

// Copy returns a new Set: a copy of s.
func (s Set) Copy() Set {
	res := Set{}
	for elem := range s {
		res[elem] = struct{}{}
	}
	return res
}

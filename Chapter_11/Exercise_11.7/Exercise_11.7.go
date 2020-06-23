package main

import (
	"bytes"
	"fmt"
	"sort"
)

type MapSet struct {
	intmap map[int]struct{}
}

func NewMapSet() *MapSet {
	return &MapSet{intmap: make(map[int]struct{})}
}

func (s *MapSet) Has(x int) bool {
	_, ok := s.intmap[x]
	return ok
}

func (s *MapSet) Add(x int) {
	s.intmap[x] = struct{}{}
}

func (s *MapSet) UnionWith(t *MapSet) {
	for k := range t.intmap {
		s.intmap[k] = struct{}{}
	}
}

func (s *MapSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for _, e := range s.Elems() {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", e)
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *MapSet) Len() int {
	return len(s.intmap)
}

func (s *MapSet) Remove(x int) {
	delete(s.intmap, x)
}

func (s *MapSet) Clear() {
	s.intmap = make(map[int]struct{})
}

func (s *MapSet) Copy() *MapSet {
	out := NewMapSet()
	for k := range s.intmap {
		out.intmap[k] = struct{}{}
	}
	return out
}

func (s *MapSet) AddAll(vals ...int) {
	for _, x := range vals {
		s.intmap[x] = struct{}{}
	}
}

func (s *MapSet) IntersectWith(t *MapSet) {
	for k := range s.intmap {
		_, ok := t.intmap[k]
		if !ok {
			delete(s.intmap, k)
		}
	}
}

func (s *MapSet) DifferenceWith(t *MapSet) {
	for k := range s.intmap {
		_, ok := t.intmap[k]
		if ok {
			delete(s.intmap, k)
		}
	}
}

func (s *MapSet) SymmetricDifference(t *MapSet) {
	for k := range t.intmap {
		_, ok := s.intmap[k]
		if ok {
			delete(s.intmap, k)
		} else {
			s.intmap[k] = struct{}{}
		}
	}
}

func (s *MapSet) Elems() []int {
	out := []int{}
	for k := range s.intmap {
		out = append(out, k)
	}
	sort.Ints(out)
	return out
}

type IntSet struct {
	words []uint64
}

func NewIntSet() *IntSet {
	return &IntSet{words: []uint64{}}
}

func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	count := 0
	for _, w := range s.words {
		for w != 0 {
			w &= w - 1
			count++
		}
	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word >= len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = []uint64{}
}

func (s *IntSet) Copy() *IntSet {
	out := &IntSet{}
	for _, w := range s.words {
		out.words = append(out.words, w)
	}
	return out
}

func (s *IntSet) AddAll(vals ...int) {
	for _, x := range vals {
		word, bit := x/64, uint(x%64)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		}
	}
}

func (s *IntSet) Elems() []int {
	out := []int{}
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				out = append(out, 64*i+j)
			}
		}
	}
	return out
}

func main() {

}

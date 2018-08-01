package main

import (
	"bytes"
	"fmt"
)

const wordLength int = 32 << (^uint(0) >> 63)

// PopCount returns the population count using the "shift off rightmost set bit and check" method
func PopCount(x uint) int {
	var result uint
	for x&(x-1) != x {
		result++
		x = x & (x - 1)
	}
	return int(result)
}

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/wordLength, uint(x%wordLength)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/wordLength, uint(x%wordLength)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWith sets s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			return
		}
	}
}

// DifferenceWith sets s to the difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= ^tword
		} else {
			return
		}
	}
}

// SymmetricDifferenceWith sets s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordLength; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", wordLength*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len returns the number of elements in the set
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += PopCount(word)
	}
	return count
}

// Remove removes an item x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/wordLength, uint(x%wordLength)
	if word >= len(s.words) {
		return
	}

	s.words[word] &= ^(1 << bit)
}

// Clear removes all elements from the set
func (s *IntSet) Clear() {
	s.words = make([]uint, 0)
}

// Copy returns a copy of the set
func (s *IntSet) Copy() *IntSet {
	copy := &IntSet{}
	for _, word := range s.words {
		copy.words = append(copy.words, word)
	}
	return copy
}

// AddAll allows multiple non-negative integers to be added at once
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		s.Add(val)
	}
}

// Elems returns a slice of the elements of the set, suitable for iteration with range
func (s *IntSet) Elems() []int {
	var values []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordLength; j++ {
			if word&(1<<uint(j)) != 0 {
				values = append(values, wordLength*i+j)
			}
		}
	}

	return values
}

func main() {
	// This main function represents extremely unstructured testing I used to verify the above
	// work to myself during development.
	x := &IntSet{}

	x.Add(1)
	x.Add(20)
	x.Add(14)
	x.Add(5000)
	fmt.Println(x)
	fmt.Println(x.Len())
	x.Remove(150)
	fmt.Println(x)
	y := x.Copy()
	x.Clear()
	fmt.Println(x)
	fmt.Println(y)
	x.AddAll(1, 2, 3, 5, 8, 13, 21, 34)
	fmt.Println(x)

	z := &IntSet{}
	z.AddAll(1, 2, 3, 5, 7, 13)

	x.SymmetricDifferenceWith(z)
	fmt.Println(x)
	fmt.Println(x.Elems())
}

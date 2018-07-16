// Exercise 4.3: Rewrite reverse to use an array pointer instead of a slice.
// Exercise 4.4: Write a version of rotate that operates in a single pass.
// Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a []string slice.
// Exercise 4.6: Write an in-place function that squashes each run of adjacent Unicode spaces
// 			     (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.

package main

import "fmt"

// reverse reverses an array of 32 ints in place. (dumb)
func reverse(s *[32]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	ary := [32]int{1, 2, 3, 4, 5, 0, 0, 3}
	reverse(&ary)
	fmt.Printf("%d\n", ary)
	reverse(&ary)
	fmt.Printf("%d\n", ary)
	s := ary[:]
	s = rotate(s, 2)
	fmt.Printf("%d\n", s)
	s = rotate(s, 32)
	fmt.Printf("%d\n", s)

	strings := []string{"hello", "world", "today", "today", "is", "is", "is", "the", "day"}
	strings = deduplicate(strings)
	fmt.Printf("%q\n", strings)

	bytes := []byte{'h', 'e', 'l', 'l', 'o', ' ', ' ', ' ', 'w', 'o', 'r', 'l', 'd', ' ', ' ', '!'}
	bytes = squashSpaces(bytes)
	fmt.Printf("%s\n", bytes)
}

// rotate a slice left by n positions in a single pass
func rotate(s []int, n int) []int {
	if len(s) == 0 {
		return s
	}
	result := make([]int, len(s))
	for i, j := 0, n%len(result); i < len(result); i, j = i+1, (j+1)%len(result) {
		result[i] = s[j]
	}
	return result
}

func deduplicate(s []string) []string {
	if len(s) == 0 {
		return s
	}
	var prev = s[0]
	var nextInsertion = 1 // keep track of the next "unused" spot in the slice
	for i := 1; i < len(s); i++ {
		if s[i] != prev {
			s[nextInsertion] = s[i]
			prev = s[i]
			nextInsertion++
		}
	}
	return s[:nextInsertion]
}

func squashSpaces(s []byte) []byte {
	if len(s) == 0 {
		return s
	}

	runes := []rune(string(s))

	followingSpace := false
	nextSlot := 0
	for i := 0; i < len(runes); i++ {
		if followingSpace && runes[i] == ' ' {

		} else {
			runes[nextSlot] = runes[i]
			followingSpace = runes[i] == ' '
			nextSlot++
		}
	}

	return []byte(string(runes[:nextSlot]))
}

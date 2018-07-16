//Daniel Harms' solutions for exercises in The Go Programming Language

// Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256 hashes.
// (See PopCount from Section 2.6.2.)

package main

import "fmt"

func countDiff(sha1 [32]byte, sha2 [32]byte) int {
	var acc int
	for i := range sha1 {
		for j := uint32(0); j < 8; j++ {
			acc += int((sha1[i] >> j & 1) ^ (sha2[i] >> j & 1))
		}
	}
	return acc
}

func main() {
	fmt.Println(countDiff([32]byte{3}, [32]byte{1}))
}

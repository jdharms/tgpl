//Daniel Harms' solutions for exercises in The Go Programming Language

// Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256 hashes.
// (See PopCount from Section 2.6.2.)

// Exercise 4.2: Write a program that prints the SHA256 hash of its standard input by default but supports
// a command-line flag to print the SHA384 or SHA512 hash instead.

package main

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
	print(countDiff([32]byte{3}, [32]byte{1}))
}

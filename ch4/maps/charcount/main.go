// Daniel Harms' solutions to exercise 4.8 in The Go Programming Language
// Exercise 4.8: Modify charcount to count letters, digits, and so on in their Unicode
// categories, using functions like unicode.IsLetter.

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)     // counts of Unicode characters
	var utflen [utf8.UTFMax + 1]int  // count of lengths of UTF-8 encodings
	utfClass := make(map[string]int) // count of characters in different classes
	invalid := 0                     // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		if unicode.IsControl(r) {
			utfClass["control"]++
		} else if unicode.IsDigit(r) {
			utfClass["digit"]++
		} else if unicode.IsLetter(r) {
			utfClass["letter"]++
		} else if unicode.IsMark(r) {
			utfClass["mark"]++
		} else if unicode.IsNumber(r) {
			utfClass["number"]++
		} else if unicode.IsSpace(r) {
			utfClass["space"]++
		} else if unicode.IsPunct(r) {
			utfClass["punctuation"]++
		} else if unicode.IsSymbol(r) {
			utfClass["symbol"]++
		} else {
			utfClass["unknown"]++
		}

		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\nclass\tcount\n")
	for i, n := range utfClass {
		fmt.Printf("%v\t%d\n", i, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

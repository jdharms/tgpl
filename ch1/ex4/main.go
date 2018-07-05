// Daniel Harms' solution to exercise 1.4 of TGPL
// Exercise 1.4:  Modify dup2 to print the names of all files in which each duplicated line occurs.

// NB: For lines duplicated at all (inside or across files) the filename will be printed
// for each instance of the line.  The simple test files I used are included alongside this file.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts)
		f.Close()
	}

	for line, n := range counts {
		if len(n) > 1 {
			fmt.Printf("%q\t%v\n", line, n)
		}
	}
}

func countLines(f *os.File, counts map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()] = append(counts[input.Text()], f.Name())
	}
	// NOTE: ignoring potential errors from input.Err()
}

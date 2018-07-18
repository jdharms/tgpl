// Exercise 4.9: Write a program wordfreq to report the frequency of each word in an input text file.
// Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into words instead of lines.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewScanner(bufio.NewReader(os.Stdin))
	in.Split(bufio.ScanWords)
	wordCount := make(map[string]int)

	for in.Scan() {
		w := in.Text()
		wordCount[w]++
	}
	fmt.Printf("%v\n", wordCount)
}

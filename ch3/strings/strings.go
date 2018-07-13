// Exercise 3.10: Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.

// Exercise 3.11: Enhance comma so that it deals correctly with floating-point numbers and an optional sign.

// Exercise 3.12: Write a function that reports whether two strings are anagrams of each other, that is, they contain the same letters in a different order.

package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// Note: recursive comma() included for reference
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	runes := []rune(s)
	n := len(runes)

	if n <= 3 {
		return s
	}

	return comma(string(runes[:n-3])) + "," + string(runes[n-3:])
}

func trimLeftChar(s string) string {
	for i := range s { // i is the index of the first byte of the string's code points
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

func enhancedComma(s string) string {
	if strings.HasPrefix(s, "+") {
		return "+" + enhancedComma(trimLeftChar(s))
	} else if strings.HasPrefix(s, "-") {
		return "-" + enhancedComma(trimLeftChar(s))
	}

	if strings.Contains(s, ".") {
		splitString := strings.Split(s, ".")
		return enhancedComma(splitString[0]) + "." + splitString[1]
	}

	n := len(s)
	if n <= 3 {
		return s
	}
	return enhancedComma(s[:n-3]) + "," + s[n-3:]
}

func nonRecursiveComma(s string) string {
	var buf bytes.Buffer
	reversed := reverse(s)
	for i, v := range reversed {
		if i > 0 && i%3 == 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune(v)
	}
	return reverse(buf.String())
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func mapString(s string) map[rune]int {
	rmap := make(map[rune]int)
	runes := []rune(s)
	for _, r := range runes {
		rmap[r]++
	}
	return rmap
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := mapString(s)
	tMap := mapString(t)
	return reflect.DeepEqual(sMap, tMap)
}

func main() {
	fmt.Println(nonRecursiveComma("12345"))
	fmt.Println(nonRecursiveComma("123456"))
	fmt.Println(nonRecursiveComma("123"))
	fmt.Println(nonRecursiveComma("1"))

	fmt.Println(enhancedComma("1234"))
	fmt.Println(enhancedComma("+1234"))
	fmt.Println(enhancedComma("-12345.4556"))
	fmt.Println(enhancedComma("-12345678345634556.99999999"))
	fmt.Println(enhancedComma("12"))

	fmt.Println(comma("プローグラム"))
	fmt.Println(isAnagram("Hello", "olleH"))
	fmt.Println(isAnagram("Hello", "Hello!"))
	fmt.Println(isAnagram("H", "H"))
}

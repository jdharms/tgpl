package main

import (
	"fmt"
)

func noReturn(val int) {
	panic(val)
}

func main() {
	defer func() {
		num := recover()
		fmt.Printf("Returned: %v\n", num)
	}()
	noReturn(5)
}

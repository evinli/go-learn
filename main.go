package main

import (
	"fmt"
)

func multiply(x, y int) int {
	return x * y
}

func main() {
	fmt.Printf("Hello, World!\n")
	fmt.Printf("2 * 3 = %d\n", multiply(2, 3))
}
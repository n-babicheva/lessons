package main

import (
	"errors"
	"fmt"
)

// add +
func add(a int, b int) int {
	return a + b
}

// sub-
func sub(a int, b int) int {
	return a - b
}

// mul*
func mul(a int, b int) int {
	return a * b
}

// div /
func div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("div is incorrect")
	}
	return a / b, nil
}
func main() {
	fmt.Println(add(1, 2))
	fmt.Println(sub(7, 4))
	fmt.Println(mul(8, 6))
	fmt.Println(div(49, 0))
}

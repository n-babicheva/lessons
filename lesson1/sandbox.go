package main

import "fmt"

func main() {
	var a int
	var b int

	fmt.Scan(&a, &b)

	fmt.Println(a + b)
	if a+b > 7 {
		fmt.Println(a - b)
	} else {
		fmt.Println("Max")
	}
}

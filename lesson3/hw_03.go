package main

import "fmt"

func main() {
	var (
		action string
		a, b   int
	)
	_, _ = fmt.Scan(&action, &a, &b)
	switch action {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Error: b mustn't be zero")
		} else {
			fmt.Println(a / b)
		}
	default:
		fmt.Println("Error: enter correct data")
	}
}

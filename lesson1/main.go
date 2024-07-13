package main

import "fmt"

func main() {
	// stdout
	fmt.Println("Hello World")

	// stdin
	var a int
	fmt.Scan(&a)

	// stdout
	fmt.Println(a * 2)
	fmt.Printf("%d\n", 3)
	fmt.Printf("%f\n", 1.5)
	fmt.Printf("%.2f\n", 1.5)
	fmt.Printf("%s\n", "hello world")
}

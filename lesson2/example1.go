package main

import (
	"fmt"
)

// os.Stdin, os.Stdout;
// errors, panic;
// Работа с пакетом strings и strconv;
func main() {
	var a int
	var b int

	fmt.Print("Input a: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Print("Input b: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Result: %d", a+b) //вывод числа
}

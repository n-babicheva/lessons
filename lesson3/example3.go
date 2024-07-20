package main

import "fmt"

func main() {
	var cond bool
	fmt.Print("Input cond: ")

	// Тут есть ":" в ":=", т.к. переменная объявлена впервые
	_, err := fmt.Scan(&cond)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// if/else
	if cond {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// if
	if cond {
		fmt.Println("true")
	}

	var a string
	fmt.Print("Input a: ")

	// Тут нет ":" в "=", т.к. переменная уже объявлена выше
	_, err = fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// if/elseif
	if a == "a" {
		fmt.Println("a")
	} else if a == "b" {
		fmt.Println("b")
	}

	// if/elseif/else
	if a == "a" {
		fmt.Println("a")
	} else if a == "b" {
		fmt.Println("b")
	} else if a == "c" {
		fmt.Println("c")
	} else {
		fmt.Println("other")
	}

	// switch/case/default
	switch a {
	case "a": // ==
		fmt.Println("a")
	case "b": // ==
		fmt.Println("b")
	case "c": // ==
		fmt.Println("c")
	default:
		fmt.Println("other")
	}
}

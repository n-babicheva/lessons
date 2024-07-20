package main

import "fmt"

func main() {
	var a int
	var err error

	_, err = fmt.Scan(&a)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(a)
	}
}

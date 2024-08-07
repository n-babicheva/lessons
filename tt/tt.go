package main

import "fmt"

var y = 11

func main() {
	y := 12

	{
		y = 15
	}

	fmt.Println(y)
}

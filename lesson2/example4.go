package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	fmt.Println("Input: ", scanner.Text())

	a, err := strconv.ParseInt(scanner.Text(), 10, 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("a: ", a*2)
}

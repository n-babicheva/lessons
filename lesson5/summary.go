package main

import (
	"errors"
	"fmt"
)

func summary(str string, count int) string {
	if count > len(str) {
		return str
	}

	return str[0:count] + "..."
}

func summary2(str string, count int) (string, error) {
	if count > len(str) {
		return "", errors.New("get count very big")
	}

	return str[0:count] + "...", nil
}

func summary3(str string, count int) string {
	result, err := summary2(str, count)
	if err != nil {
		return ""
	}

	return result
}

func main() {
	str := "hello world"

	fmt.Println(summary(str, 10))

	fmt.Println()

	fmt.Println(summary2(str, 10))
	fmt.Println(summary2(str, 12))

	fmt.Println()

	fmt.Println(summary3(str, 10))
	fmt.Println(summary3(str, 12))
}

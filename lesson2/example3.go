package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string
	fmt.Println(s1)

	s2 := "hello world"
	fmt.Println(s2)

	fmt.Println(len(s2))

	/**
		hello world
		0123456789X
	   [0...4]
	*/

	fmt.Println(s2[:5]) // hello
	fmt.Println(s2[6:]) // world

	fmt.Println(strings.Contains("hello world", "world")) // true

	fmt.Println(strings.Count("world world", "world")) // 2

	fmt.Println(strings.Index("hello world", "world"))

	str := `asdfasdfasdf
asdfasdfasdfsadf hello world
adfasdfasdf
`

	pos := strings.Index(str, "hello")

	fmt.Println(str[pos : pos+11])

	fmt.Println(strings.ToUpper("hello world")) // HELLO WORLD
	fmt.Println(strings.ToLower("HELLO WORLD")) // hello world

	fmt.Println(strings.Repeat("hello world ", 5))

	fmt.Println(strings.Fields("hola mi")) // ["hola", "mi"]
}

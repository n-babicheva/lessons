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

	fmt.Println(s2[:5]) // отрезать от строки первые 5 символов
	fmt.Println(s2[6:]) // отрезаем от строки все символы начиная с 6го

	fmt.Println(strings.Contains("hello world", "world")) // провека вхождения подстроки в строку - false/true

	fmt.Println(strings.Count("world world", "world")) // подсчет количества вхождений подстроки в строку

	fmt.Println(strings.Index("hello world", "world")) // индекс начала вхождения

	str := `asdfasdfasdf
	asdfasdfasdfsadf hello world
	adfasdfasdf
	`

	pos := strings.Index(str, "hello")

	fmt.Println(str[pos : pos+11]) //вырезать из строки подстроку

	fmt.Println(strings.ToUpper("hello world")) // HELLO WORLD
	fmt.Println(strings.ToLower("HELLO WORLD")) // hello world

	fmt.Println(strings.Repeat("hello world ", 5))

	fmt.Println(strings.Fields("hola mi")) // ["hola", "mi"]
}

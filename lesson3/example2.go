package main

import "fmt"

func main() {
	var b bool

	fmt.Println(b)

	fmt.Printf("hello => %v\n", []byte("hello"))
	fmt.Printf("hello => %v\n", []rune("hello"))

	fmt.Printf("привет => %v\n", []byte("привет"))
	fmt.Printf("привет => %v\n", []rune("привет"))

	// type data interface | any
	var d any = "hello world"
	var e any = 1 // int
	var f interface{} = true

	fmt.Println(d, e, f)

	o1, ok := e.(int) // 1, true - перевод переменной в нужный тип
	fmt.Printf("%+v %+v\n", ok, o1)
	o2, ok := e.(string) // "", false - перевод переменной в нужный тип
	fmt.Printf("%+v %+v\n", ok, o2)
}

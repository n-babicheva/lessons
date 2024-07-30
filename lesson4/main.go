package main

import "fmt"

func main() {
	// бесконечный цикл
	//for {
	//	fmt.Println("Hello World")
	//
	//	time.Sleep(1 * time.Second)
	//}

	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}

	//for i := range 10 {
	//	fmt.Println(i)
	//}

	//sum := 0
	//for i := 0; i < 10; i++ {
	//	if i == 2 {
	//		continue
	//	}
	//
	//	sum += i
	//}
	//
	//fmt.Println(sum)

	//sum := 0
	//for i := 0; i < 10; i++ {
	//	if i == 2 {
	//		break
	//	}
	//
	//	// sum=0, i=0. newSum=0
	//	// sum=0, i=1. newSum=1
	//	// sum=1, i=2. break
	//	sum += i
	//}
	//
	//fmt.Println(sum)

	//var a int
	//// [-inf..-1, 1..100]
	//for i := 0; a == 0 || a > 100; i++ {
	//	fmt.Scan(&a)
	//
	//	fmt.Println("a=", a)
	//}
	//
	//fmt.Println("result=", a)

	//// {0: 1, 1: 2, 3: 4, ... 9: 10}
	//slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//for key, value := range slice {
	//	fmt.Println("key=", key, "value=", value)
	//}
	//
	//fmt.Println("----")
	//
	//slice = []int{-55, 13, 0, 66, 99}
	//for key, value := range slice {
	//	fmt.Println("key=", key, "value=", value)
	//}
	//
	//fmt.Println("----")

	a := make([]int, 0)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 1)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 2)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 3)
	fmt.Println(a, len(a), cap(a))

	a = append(a, 4, 5)
	fmt.Println(a, len(a), cap(a))

	for k, v := range a {
		fmt.Println(k, v, a[k])
	}
}

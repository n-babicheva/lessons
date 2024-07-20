package main

import (
	"fmt"
)

// os.Stdin, os.Stdout;
// errors, panic;
// Работа с пакетом strings и strconv;
func main() {
	// объявляем переменные a и b типа int
	var a int
	var b int

	// запрашиваем у пользователя ввод числа a
	// считываем число из стандартного ввода и записываем его в переменную a,
	// также проверяем на ошибку и выводим ее, если она есть
	fmt.Print("Input a: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// запрашиваем у пользователя ввод числа b
	// считываем число из стандартного ввода и записываем его в переменную b,
	// также проверяем на ошибку и выводим ее, если она есть
	fmt.Print("Input b: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// выводим результат сложения чисел a и b
	fmt.Printf("Result: %d", a+b)
}

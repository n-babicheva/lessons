package main

import (
	"fmt"
	"strings"
)

// Пример работы с пакетом strings
func main() {
	// Объявляем переменную s1 типа string, инициализируем ее пустой строкой
	var s1 string
	fmt.Println(s1) // Выводим пустую строку

	// Объявляем переменную s2 типа string, инициализируем ее строкой "hello world"
	s2 := "hello world"
	fmt.Println(s2) // Выводим строку "hello world"

	// Выводим длину строки s2
	fmt.Println(len(s2)) // Выводим 11

	// Выводим первые 5 символов строки s2
	fmt.Println(s2[:5]) // Выводим строку "hello"

	// Выводим все символы строки s2 начиная с 6го
	fmt.Println(s2[6:]) // Выводим строку "world"

	// Проверяем, содержит ли строка "hello world" подстроку "world"
	fmt.Println(strings.Contains("hello world", "world")) // Выводим true

	// Подсчитываем количество вхождений подстроки "world" в строку "world world"
	fmt.Println(strings.Count("world world", "world")) // Выводим число 2

	// Выводим индекс начала вхождения подстроки "world" в строку "hello world"
	fmt.Println(strings.Index("hello world", "world")) // Выводим число 6

	// Объявляем строку str
	str := `asdfasdfasdf
	asdfasdfasdfsadf hello world
	adfasdfasdf
	`

	// Находим индекс начала вхождения подстроки "hello" в строку str
	pos := strings.Index(str, "hello")

	// Выводим подстроку, которая начинается с индекса pos и длиной 11 символов
	fmt.Println(str[pos : pos+11]) // Выводим строку "hello world"

	// Выводим строку "hello world" в верхнем регистре
	fmt.Println(strings.ToUpper("hello world")) // Выводим строку "HELLO WORLD"

	// Выводим строку "HELLO WORLD" в нижнем регистре
	fmt.Println(strings.ToLower("HELLO WORLD")) // Выводим строку "hello world"

	// Повторяем строку "hello world" 5 раз и выводим результат
	fmt.Println(strings.Repeat("hello world ", 5)) // Выводим строку "hello world hello world hello world hello world hello world "

	// Разбиваем строку "hola mi" на массив строк по пробелам
	fmt.Println(strings.Fields("hola mi")) // Выводим массив строк ["hola", "mi"]
}

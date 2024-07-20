package main

import "fmt"

func main() {
	var cond bool
	fmt.Print("Input cond: ")

	// Тут есть ":" в ":=", т.к. переменная объявлена впервые
	// fmt.Scan() - считывает данные из стандартного ввода
	// fmt.Scan() - возвращает 2 значения: количество успешно считанных
	// значений и ошибку, если она произошла
	_, err := fmt.Scan(&cond)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// if/else
	// if - проверяет условие, если оно true, то выполняется код
	// в теле if, если false, то выполняется код в теле else
	if cond {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	// if
	// if - тоже самое, что и if/else, но если условие false,
	// то ничего не происходит
	if cond {
		fmt.Println("true")
	}

	var a string
	fmt.Print("Input a: ")

	// Тут нет ":" в "=", т.к. переменная уже объявлена выше
	// fmt.Scan() - считывает данные из стандартного ввода
	// fmt.Scan() - возвращает 2 значения: количество успешно считанных
	// значений и ошибку, если она произошла
	_, err = fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// if/elseif
	// if - проверяет условие, если оно true, то выполняется код
	// в теле if, если false, то выполняется код в теле else
	// elseif - проверяет условие, если оно true, то выполняется код
	// в теле elseif, если false, то выполняется код в теле else
	if a == "a" {
		fmt.Println("a")
	} else if a == "b" {
		fmt.Println("b")
	}

	// if/elseif/else
	// if - проверяет условие, если оно true, то выполняется код
	// в теле if, если false, то выполняется код в теле else
	// elseif - проверяет условие, если оно true, то выполняется код
	// в теле elseif, если false, то выполняется код в теле else
	// else - выполняется код, если все предыдущие условия false
	if a == "a" {
		fmt.Println("a")
	} else if a == "b" {
		fmt.Println("b")
	} else if a == "c" {
		fmt.Println("c")
	} else {
		fmt.Println("other")
	}

	// switch/case/default
	// switch - проверяет значение выражения, которое передается
	// в switch, и сравнивает его со всеми значениями case
	// default - выполняется код, если не было найдено
	// ни одного значения case, которое равно значению выражения
	switch a {
	case "a": // ==
		fmt.Println("a")
	case "b": // ==
		fmt.Println("b")
	case "c": // ==
		fmt.Println("c")
	default:
		fmt.Println("other")
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()                            //сканируем ввод до enter
	pos := strings.Index(scanner.Text(), " ") // index - ищет ПЕРВОЕ ВХОЖДЕНИЕ подстроки "пробел"
	fmt.Println(scanner.Text()[:pos])         // str[x : y]. x - откуда, y - сколько
	fmt.Println("Input: ", scanner.Text())    //вывод отсканированного значения
}

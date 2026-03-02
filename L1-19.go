package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Разработать программу, которая переворачивает подаваемую на вход строку.

// Например: при вводе строки «главрыба» вывод должен быть «абырвалг».

// Учтите, что символы могут быть в Unicode (русские буквы, emoji и пр.),
// то есть просто iterating по байтам может не подойти — нужен срез рун ([]rune).

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите строку")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()

		if strings.ToLower(input) == "выход" {
			fmt.Println("Завершение работы...")
			break
		}

		result := reverse(input)
		fmt.Println("До: ", input)
		fmt.Println("После: ", result)
	}

}

func reverse(str string) string {
	runes := []rune(str)
	left := 0
	right := len(runes) - 1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

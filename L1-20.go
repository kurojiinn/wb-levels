package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Разработать программу, которая переворачивает порядок слов в строке.

// Пример: входная строка:

// «snow dog sun», выход: «sun dog snow».

// Считайте, что слова разделяются одиночным пробелом. Постарайтесь не использовать дополнительные срезы, а выполнять операцию «на месте».
//

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

		runes := []rune(input)
		reverse(runes)

		start := 0

		//логика для перевората слов в предложении
		for i := 0; i < len(runes); i++ {
			if runes[i] == ' ' {
				reverse(runes[start:i])
				start = i + 1
			}
		}
		reverse(runes[start:])

		fmt.Println("Результат: ", string(runes))
	}
}

func reverse(runes []rune) {
	left := 0
	right := len(runes) - 1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
}

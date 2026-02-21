package main

import (
	"fmt"
)

//Разработать программу, которая в runtime способна определить тип переменной, переданной в неё (на вход подаётся interface{}).
// Типы, которые нужно распознавать: int, string, bool, chan (канал).

// Подсказка: оператор типа switch v.(type) поможет в решении.

func main() {
	ch := make(chan int)
	worker(ch)
}

func worker(variable interface{}) {
	switch v := variable.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		fmt.Println("chan", v)
	}
}

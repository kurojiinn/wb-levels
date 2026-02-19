package main

import "fmt"

//Поменять местами два числа без использования временной переменной.

// Подсказка: примените сложение/вычитание или XOR-обмен.
//

func main() {
	withSumMin()
}

func regular() {
	a := 10
	b := 20

	a, b = b, a

	fmt.Printf("a: %v, b: %v", a, b)
}

func withSumMin() {
	a := 10
	b := 20

	a += b    //a = 30 b = 20
	b = a - b // b=30 - 20 = 10
	a = a - b //a = 30 - 10 = 20

	fmt.Printf("a: %v, b: %v", a, b)
}

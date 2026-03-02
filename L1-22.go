package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a, b, значения которых > 2^20 (больше 1 миллион).

// Комментарий: в Go тип int справится с такими числами, но обратите
// внимание на возможное переполнение для ещё больших значений. Для очень больших чисел можно использовать math/big.
//

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Введите первое число: ")
		if !scanner.Scan() {
			break
		}
		inputA := strings.TrimSpace(scanner.Text())
		if inputA == "exit" {
			fmt.Println("Завершение работы")
			break
		}
		fmt.Println("Введите второе число: ")
		if !scanner.Scan() {
			break
		}
		inputB := strings.TrimSpace(scanner.Text())

		a, ok1 := new(big.Int).SetString(inputA, 10)
		b, ok2 := new(big.Int).SetString(inputB, 10)

		if !ok1 || !ok2 {
			fmt.Println("ошибка преобразования")
			continue
		}

		calculate(a, b)
	}

}

func calculate(a, b *big.Int) {
	//Сложение
	resAdd := new(big.Int).Add(a, b)
	fmt.Printf("Сумма: %s\n", resAdd.String())
	// Вычитание
	resSub := new(big.Int).Sub(a, b)
	fmt.Printf("Разность: %s\n", resSub.String())
	// Умножение
	resMul := new(big.Int).Mul(a, b)
	fmt.Printf("Произведение: %s\n", resMul.String())
	// Деление
	if b.Sign() == 0 {
		fmt.Println("Ошибка: нельзя делить на ноль")
	} else {
		resDev := new(big.Int).Div(a, b)
		fmt.Printf("Деление: %s\n", resDev.String())
	}
}

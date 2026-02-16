package main

import (
	"fmt"
)

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа x из массива, во второй – результат операции x*2.
// После этого данные из второго канала должны выводиться в stdout. То есть, организуйте конвейер из двух этапов с горутинами:
// генерация чисел и их обработка.
// Убедитесь, что чтение из второго канала корректно завершается.

func generate(ch1 chan int, arr []int) {
	for _, val := range arr {
		ch1 <- val
	}
	defer close(ch1)
}

func worker(ch chan int, result chan int) {
	for val := range ch {
		result <- val * 2
	}
	defer close(result)
}

func main() {
	ch1 := make(chan int)
	res := make(chan int)
	arr := []int{1, 2, 3, 4, 5}

	go generate(ch1, arr)
	go worker(ch1, res)

	for val := range res {
		fmt.Println(val)
	}
}

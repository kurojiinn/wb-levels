package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значения квадратов чисел, взятых из массива [2,4,6,8,10], и выведет результаты в stdout.
//
//Подсказка: запусти несколько горутин, каждая из которых возводит число в квадрат.

func main() {
	wg := &sync.WaitGroup{}
	arr := []int{2, 4, 6, 8, 10}
	result := make(chan int, len(arr))

	square(arr, wg,result)

	go func () {
		wg.Wait()
		close(result)
	} ()

	for val := range result {
		fmt.Println(val)
	}

}

func square(arr []int, wg *sync.WaitGroup, result chan<- int) {
	for i := 0; i < len(arr); i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			result <- val * val
		}(arr[i])
	}
}

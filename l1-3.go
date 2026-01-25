package main

import (
	"fmt"
	"os/signal"
	"os"
	"sync"
)


//Реализовать набор из N воркеров, которые читают данные из этого канала и выводят их в stdout.
func worker(jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range jobs{
		fmt.Println(val)
	}
}
func main () {
	var workerCount int
	jobs := make(chan int)

	wg := &sync.WaitGroup{}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	//Программа должна принимать параметром количество воркеров
	fmt.Println("введите кол-во воркеров")
	_, err := fmt.Scan(&workerCount)
	if err != nil {
		return
	}

	if workerCount <=0 {
		fmt.Println("Введено некорректное кол-во воркеров")
		return
	}
	//и при старте создавать указанное число горутин-воркеров.
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(jobs, wg)
	}

	//Реализовать постоянную запись данных в канал (в главной горутине).
	Flag:
		for i := 0; ;i++{
			select {
				case jobs <- i:
				case <-sigChan:
					fmt.Println("выходим из цикла")
					break Flag
			}
		}

	close(jobs)
	fmt.Println("Закрыли канал")
	wg.Wait()
}

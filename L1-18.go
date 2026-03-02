package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// //Конкурентный счетчик
// Реализовать структуру-счётчик, которая будет инкрементироваться в конкурентной среде (т.е. из нескольких горутин).
//  По завершению программы структура должна выводить итоговое значение счётчика.

// Подсказка: вам понадобится механизм синхронизации, например, sync.Mutex или sync/Atomic для безопасного инкремента.
const workerNums = 5

type CounterWithMutex struct {
	mu    sync.Mutex
	value int
}

func (c *CounterWithMutex) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func main() {
	wg := &sync.WaitGroup{}
	counterMu := &CounterWithMutex{}
	for range workerNums {
		go counterMu.increment()
	}

	//через атомики
	var aint atomic.Int64

	for range workerNums {
		wg.Add(1)
		go func() {
			defer wg.Done()
			aint.Add(1)
		}()
	}

	wg.Wait()
	fmt.Println("Результат с мьютексом: ", counterMu.value)
	fmt.Println("Результат: ", aint.Load())
}

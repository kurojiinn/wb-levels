package main

import (
	"fmt"
	"sync"
)

//Реализовать безопасную для конкуренции запись данных в структуру map.

// Подсказка: необходимость использования синхронизации (например, sync.Mutex или встроенная concurrent-map).

// Проверьте работу кода на гонки (util go run -race).
func main() {
	UsingBasicMap()
	UsingSyncMap()
}

func UsingBasicMap() {
	mp := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.RWMutex{}

	for i := range 5 {
		wg.Add(1)
		val := i
		go func() {
			defer wg.Done()
			mu.Lock()
			mp[val] = val + 10
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("Результат:", mp)
}

func UsingSyncMap() {
	mp := sync.Map{}
	wg := sync.WaitGroup{}

	for i := range 5 {
		wg.Add(1)
		go func(val int) {
			mp.Store(val, val+10)
		}(i)
	}

	wg.Wait()

	value, _ := mp.Load(0)
	fmt.Println(value)
}

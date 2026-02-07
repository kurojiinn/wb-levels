package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

// Классические подходы: выход по условию, через канал уведомления,
// через контекст, прекращение работы runtime.Goexit() и др.

// Продемонстрируйте каждый способ в отдельном фрагменте кода.
func main() {
	data := make(chan string)
	wg := &sync.WaitGroup{}
	ticker := time.NewTicker(500 * time.Millisecond)
	aflag := &atomic.Bool{}
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan struct{})

	wg.Add(5)
	go func() {
		defer wg.Done()
		CancelWithContext(ctx, ticker)

	}()

	go func() {
		defer wg.Done()
		CancelWithChanelNotify(ch, ticker)
	}()

	go func() {
		defer wg.Done()
		CancelWithConditional(aflag)
	}()

	go func() {
		defer wg.Done()
		CancelWithGoexit()
	}()

	go func() {
		defer wg.Done()
		CancelWithClosedChannel(data)
	}()

	time.Sleep(2 * time.Second)

	cancel()
	close(ch)
	close(data)
	aflag.Store(true)

	wg.Wait()

}

// через контекст
func CancelWithContext(ctx context.Context, ticker *time.Ticker) {
	for {
		select {
		case <-ticker.C:
			fmt.Println("25ms")
		case <-ctx.Done():
			return
		}
	}
}

// прекращение работы runtime.Goexit()
func CancelWithGoexit() {
	fmt.Println("Что то делаем")
	defer fmt.Println("Сначала испольнится defer")
	runtime.Goexit()
}
func CancelWithConditional(flag *atomic.Bool) { //выход по условию
	for {
		if flag.Load() {
			return
		}
	}
}

// через канал уведомления
func CancelWithChanelNotify(ch <-chan struct{}, ticker *time.Ticker) {
	for {
		select {
		case <-ch:
			return
		case <-ticker.C:
			fmt.Println("25")
		}
	}
}

// завершение по закрытию рабочего канала
func CancelWithClosedChannel(data <-chan string) {
	for val := range data {
		fmt.Println(val)
	}
}

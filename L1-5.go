package main

import (
	"context"
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала – читать эти значения. По истечении N секунд программа должна завершаться.

// Подсказка: используйте time.After или таймер для ограничения времени работы.
func send(ctx context.Context, ch chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case ch <- 1:
		}
	}
}

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		send(ctx, ch)
	}()
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			cancel()
			return
		case val := <-ch:
			fmt.Println(val)
		}
	}
}

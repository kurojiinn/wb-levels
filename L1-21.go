package main

import (
	"fmt"
)

//Реализовать паттерн проектирования «Адаптер» на любом примере.

// Описание: паттерн Adapter позволяет сконвертировать интерфейс одного класса в интерфейс другого, который ожидает клиент.

// Продемонстрируйте на простом примере в Go: у вас есть существующий интерфейс (или структура) и другой, несовместимый по
// интерфейсу потребитель — напишите адаптер, который реализует нужный интерфейс и делегирует вызовы к встроенному объекту.

// Поясните применимость паттерна, его плюсы и минусы, а также приведите реальные примеры использования.

type Logger interface {
	Log(message string)
}

type SimpleLogger struct {
}

type ZapLogger struct{}

type ZapLoggerAdapter struct {
	zapLogger *ZapLogger //храним несовместимый объект
}

// он уже будет реализовать метод Log, чтобы удовлетворить интерфейсу Logger
func (zla *ZapLoggerAdapter) Log(message string) {
	//тут уже будем вызывать несовместимый метод, который выполняет туже логику
	zla.zapLogger.Info(message)
}

func (z *ZapLogger) Info(message string) {
	fmt.Println("[ZAP]:", message)
}

func (s *SimpleLogger) Log(message string) {
	fmt.Println(message)
}

func doSomething(logger Logger) {
	logger.Log("message")
}

func main() {
	zapS := &ZapLogger{}
	adapter := ZapLoggerAdapter{zapLogger: zapS}
	doSomething(&adapter)

}

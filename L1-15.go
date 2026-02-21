package main

import "strings"

// Рассмотреть следующий код и ответить на вопросы: к каким негативным последствиям он может привести и как это исправить?

// Приведите корректный пример реализации.
// Вопрос: что происходит с переменной justString?

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}

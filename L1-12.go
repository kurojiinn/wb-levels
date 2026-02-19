package main

import "fmt"

//Имеется последовательность строк: ("cat", "cat", "dog", "cat", "tree"). Создать для неё собственное множество.

//Ожидается: получить набор уникальных слов. Для примера, множество = {"cat", "dog", "tree"}.

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(worker(arr))
}

func worker(arr []string) []string {
	unique := make(map[string]struct{})

	for _, val := range arr {
		unique[val] = struct{}{}
	}

	res := make([]string, 0, len(unique))
	for k, _ := range unique {
		res = append(res, k)
	}

	return res
}

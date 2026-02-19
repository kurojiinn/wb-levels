package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств (например, двух слайсов) — т.е. вывести элементы, присутствующие и в первом, и во втором.

// Пример:
// A = {1,2,3}
// B = {2,3,4}
// Пересечение = {2,3}
//
//

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	fmt.Println(worker(A, B))
}

func worker(A []int, B []int) []int {
	metch := make(map[int]struct{})
	res := []int{}

	for _, val := range A {
		metch[val] = struct{}{}
	}

	for _, val := range B {
		if _, ok := metch[val]; ok {
			res = append(res, val)
		}
	}

	return res
}

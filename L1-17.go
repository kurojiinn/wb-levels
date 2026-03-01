package main

func main() {

}

//Реализовать алгоритм бинарного поиска встроенными методами языка.
// Функция должна принимать отсортированный слайс и искомый элемент,
// возвращать индекс элемента или -1, если элемент не найден.

func binarySearch(arr []int, target int) int {
	right := len(arr) - 1
	left := 0

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

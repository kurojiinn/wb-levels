package main

func qsort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[len(arr)/2]

	left := []int{}
	middle := []int{}
	right := []int{}

	for _, v := range arr {
		switch {
		case v < pivot:
			left = append(left, v)
		case v == pivot:
			middle = append(middle, v)
		default:
			right = append(right, v)
		}
	}

	result := append(qsort(left), middle...)
	result = append(result, qsort(right)...)
	return result
}

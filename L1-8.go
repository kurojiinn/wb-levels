package main

import "fmt"

// SetBit устанавливает i-й бит val в target (true=1, false=0)
func SetBit(val int64, i int, target bool) int64 {
	mask := int64(1 << i)
	if target {
		return val | mask
	}
	return val &^ mask
}

func main() {
	val := int64(5)
	result := SetBit(val, 1, false) // i=1 в 0 (но уже 0)
	fmt.Println(result)

	result = SetBit(val, 0, false) // i=0 в 0 → 4
	fmt.Println(result)
}

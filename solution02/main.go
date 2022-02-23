package main

import (
	"fmt"
)

const n = 12

func firstIndexOfNum(a [n]int, num int) (index int) {
	index = -1
	for i := 0; i < n; i++ {
		if a[i] == num {
			index = i
			return
		}
	}
	return
}

func main() {
	a := [n]int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(a)

	num := 2
	fmt.Printf("Индекс первого вхождения числа %v: %v\n", num, firstIndexOfNum(a, num))
}

package main

import (
	"fmt"
)

const n = 12

func firstIndexOfNum(a [n]int, num int) (index int) {
	index = -1
	min := 0
	max := n - 1

	for max >= min {
		middle := (max + min) / 2
		if a[middle] == num {
			index = middle
			break
		} else if a[middle] > num {
			max = middle - 1
		} else {
			min = middle + 1
		}
	}

	for index > 0 && index == num {
		index--
	}
	return
}

func main() {
	a := [n]int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(a)

	num := 2
	fmt.Printf("Индекс первого вхождения числа %v: %v\n", num, firstIndexOfNum(a, num))
}

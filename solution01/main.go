package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func countNumbersAfterThisNumber(a [n]int, num int) int {
	numbersAfterThisNum := 0
	for i := 0; i < n; i++ {
		if a[i] == num {
			numbersAfterThisNum = n - 1 - i
		}
	}
	return numbersAfterThisNum
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var a [n]int

	for i := 0; i < n; i++ {
		a[i] = rand.Intn(10)
		//fmt.Println(a[i])
	}

	fmt.Println(a)

	num := a[rand.Intn(n)]
	fmt.Printf("Чисел в массиве после числа %v: %v\n", num, countNumbersAfterThisNumber(a, num))

	num = n + 1
	fmt.Printf("Чисел в массиве после числа %v: %v\n", num, countNumbersAfterThisNumber(a, num))
}

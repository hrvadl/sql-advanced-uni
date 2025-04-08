package main

import (
	"fmt"
)

func toDigits(n int) []int {
	digits := []int{}
	if n == 0 {
		digits = append(digits, 0)
	} else {
		tempN := n
		for tempN > 0 {
			digit := tempN % 10
			digits = append([]int{digit}, digits...)
			tempN /= 10
		}
	}

	return digits
}

func main() {
	var n int
	fmt.Print("Введіть ціле невід'ємне число n: ")
	fmt.Scan(&n)

	digits := toDigits(n)
	for i, digit := range digits {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(digit)
	}
	fmt.Println()
}

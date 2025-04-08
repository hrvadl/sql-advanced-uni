package main

import "fmt"

func removeDigits(n int) int {
	result := 0
	multiplier := 1

	for n > 0 {
		digit := n % 10
		n /= 10

		if digit != 0 && digit != 5 {
			result += digit * multiplier
			multiplier *= 10
		}
	}

	return result
}

func main() {
	var n int
	fmt.Print("Please, input  n: ")
	fmt.Scan(&n)

	result := removeDigits(n)
	fmt.Printf("The number %d without 5 and 0: %d\n", n, result)
}

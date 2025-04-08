package main

import (
	"fmt"
	"math"
)

func findDivisors(n int) []int {
	divisors := []int{}

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			divisors = append(divisors, i)

			if i != n/i {
				divisors = append(divisors, n/i)
			}
		}
	}

	for i := 0; i < len(divisors); i++ {
		for j := i + 1; j < len(divisors); j++ {
			if divisors[i] > divisors[j] {
				divisors[i], divisors[j] = divisors[j], divisors[i]
			}
		}
	}

	return divisors
}

func main() {
	var n int
	fmt.Print("Please, input number n: ")
	fmt.Scan(&n)

	divisors := findDivisors(n)

	fmt.Printf("Divisors of number %d: ", n)
	for i, div := range divisors {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(div)
	}
	fmt.Println()
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Print("Input the number: ")
	fmt.Scan(&n)

	var (
		stringNumber = strconv.Itoa(n)
		digitCount   = len(stringNumber)
		firstDigit   = int(stringNumber[0] - '0')
	)

	sum := 0
	tempN := n
	for tempN > 0 {
		sum += tempN % 10
		tempN /= 10
	}

	fmt.Printf("Digit count in number %d: %d\n", n, digitCount)
	fmt.Printf("Digit sum of number %d: %d\n", n, sum)
	fmt.Printf("First digit of number %d: %d\n", n, firstDigit)
}

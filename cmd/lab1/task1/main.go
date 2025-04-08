package main

import (
	"fmt"
	"strconv"
)

func process(n int) (digitCount, digitSum, firstDigit int) {
	stringNumber := strconv.Itoa(n)
	digitCount = len(stringNumber)
	firstDigit = int(stringNumber[0] - '0')

	tempN := n
	for tempN > 0 {
		digitSum += tempN % 10
		tempN /= 10
	}

	return
}

func main() {
	var n int
	fmt.Print("Input the number: ")
	fmt.Scan(&n)
	digitCount, digitSum, firstDigit := process(n)

	fmt.Printf("Digit count in number %d: %d\n", n, digitCount)
	fmt.Printf("Digit sum of number %d: %d\n", n, digitSum)
	fmt.Printf("First digit of number %d: %d\n", n, firstDigit)
}

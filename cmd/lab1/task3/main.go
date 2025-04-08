package main

import "fmt"

func gcd(m, n int) int {
	if n == 0 {
		return m
	}

	return gcd(n, m%n)
}

func main() {
	var m, n int
	fmt.Print("Please, input m and n: ")
	fmt.Scan(&m, &n)

	result := gcd(m, n)
	fmt.Printf("GCD(%d, %d) = %d\n", m, n, result)
}

package main

import (
	"fmt"
	"strconv"
)

func process(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func main() {
	var n int
	fmt.Print("Please, input n: ")
	fmt.Scan(&n)

	fmt.Printf("Number %d converted to the binary: %s", n, process(n))
}

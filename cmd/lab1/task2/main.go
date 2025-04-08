package main

import (
	"fmt"
	"math"
)

func main() {
	var epsilon float64
	fmt.Print("Input precise ε (I.E: 0.001): ")
	fmt.Scan(&epsilon)

	sum := 0.0
	term := 0.0
	i := 0

	for {
		term = 1.0 / (math.Pow(4, float64(i)) + math.Pow(5, float64(i+2)))

		if math.Abs(term) < epsilon {
			break
		}

		sum += term
		i++
	}

	fmt.Printf("Sum with precise %f: %f\n", epsilon, sum)
	fmt.Printf("Count of additions: %d\n", i)
}

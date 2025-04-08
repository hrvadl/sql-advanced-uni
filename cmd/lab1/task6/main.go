package main

import "fmt"

func process() []int {
	res := []int{}
	for num := 1000; num <= 9999; num++ {
		lastDigit := num % 10
		if lastDigit != 0 && num%23 == 0 && num%lastDigit == 0 {
			res = append(res, num)
		}
	}

	return res
}

func main() {
	numbers := process()

	fmt.Printf(
		"Кількість чотиризначних чисел, що діляться на 23 та на свою останню цифру: %d\n",
		len(numbers),
	)

	for _, num := range numbers {
		fmt.Printf("%d (ділиться на %d і на останню цифру)\n", num, 23)
	}
}

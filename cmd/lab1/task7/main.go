package main

import "fmt"

func main() {
	baseNum := 1022

	for leftDigit := 1; leftDigit <= 9; leftDigit++ {
		for rightDigit := 0; rightDigit <= 9; rightDigit++ {
			num := leftDigit*100000 + baseNum*10 + rightDigit

			if num%7 == 0 && num%8 == 0 && num%9 == 0 {
				fmt.Printf("Знайдено число: %d\n", num)
				fmt.Printf("Ліва цифра: %d, права цифра: %d\n", leftDigit, rightDigit)

				fmt.Printf("Число %d ділиться на 7, 8, 9\n", num)
			}
		}
	}
}

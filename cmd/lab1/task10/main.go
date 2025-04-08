package main

import (
	"fmt"
)

func hanoi(n int, source, auxiliary, target int, moves *[]string) {
	if n == 1 {
		*moves = append(
			*moves,
			fmt.Sprintf("Перемістити диск 1 зі стрижня %d на стрижень %d", source, target),
		)
		return
	}

	// Крок 1: перемістити n-1 дисків з вихідного на допоміжний стрижень
	hanoi(n-1, source, target, auxiliary, moves)

	// Крок 2: перемістити найбільший диск з вихідного на цільовий стрижень
	*moves = append(
		*moves,
		fmt.Sprintf("Перемістити диск %d зі стрижня %d на стрижень %d", n, source, target),
	)

	// Крок 3: перемістити n-1 дисків з допоміжного на цільовий стрижень
	hanoi(n-1, auxiliary, source, target, moves)
}

func main() {
	var n int
	fmt.Print("Введіть кількість дисків: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Кількість дисків повинна бути додатним числом")
		return
	}

	fmt.Println("Послідовність кроків для переміщення", n, "дисків:")
	moves := []string{}
	hanoi(n, 1, 2, 3, &moves)

	for i, move := range moves {
		fmt.Printf("%d. %s\n", i+1, move)
	}

	fmt.Printf("\nВсього кроків: %d\n", len(moves))
	fmt.Printf("Найменша кількість кроків: %d\n", (1<<n)-1)
}

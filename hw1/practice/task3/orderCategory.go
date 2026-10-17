package main

import (
	"fmt"
	"os"
)

func main() {
	var amount int
	fmt.Fscan(os.Stdin, &amount)

	if amount < 0 {
		fmt.Println("Ошибка: сумма заказа не может быть отрицательной")
		return
	}

	switch {
	case amount < 10000:
		fmt.Println("Сумма заказа: ", amount, " коп.")
		fmt.Println("Категория: small")
	case amount < 50000:
		fmt.Println("Сумма заказа: ", amount, " коп.")
		fmt.Println("Категория: regular")
	default:
		fmt.Println("Сумма заказа: ", amount, " коп.")
		fmt.Println("Категория: large")
	}

}

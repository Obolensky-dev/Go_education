package main

import "fmt"

func main() {

	var a, b int
	var op string

	fmt.Print("Введите два числа через пробел: ")
	fmt.Scan(&a, &b)

	fmt.Print("Введите операцию (+, -, *, /): ")
	fmt.Scan(&op)

	switch op {
	case "+":
		fmt.Println("Результат:", a+b)
	case "-":
		fmt.Println("Результат:", a-b)
	case "*":
		fmt.Println("Результат:", a*b)
	case "/":
		if b != 0 {
			fmt.Println("Результат:", a/b)
		} else {
			fmt.Println("Ошибка: деление на ноль")
		}
	default:
		fmt.Println("Неизвестная операция")
	}

}

package main

import "fmt"

func main() {

	var numb int
	var stop string
	var modulo int

	fmt.Println("Введите число: ")
	fmt.Scanln(&numb)

	modulo = numb % 2

	switch modulo {
	case 0:
		fmt.Println("Четное")
	default:
		fmt.Println("Нечетное")
	}

	fmt.Println("Нажмите Enter для выхода...")
	fmt.Scanln(&stop)

}

package main

import "fmt"

func main() {

	var numb int
	var stop string

	fmt.Println("Введите число: ")
	fmt.Scanln(&numb)

	if numb > 0 {
		fmt.Println("Положительное")
	} else if numb == 0 {
		fmt.Println("Ноль")
	} else {
		fmt.Println("Отрицательное")
	}

	fmt.Println("Нажмите Enter для выхода...")
	fmt.Scanln(&stop)

}

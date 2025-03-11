package main

import "fmt"

func main() {

	var numb1 int
	var numb2 int
	var operation int

	fmt.Println("Введите первое число: ")
	fmt.Scanln(&numb1)
	fmt.Println("Введите второе число: ")
	fmt.Scanln(&numb2)
	fmt.Printf("Выберите операцию и наберите соответствующее число:\n0. +\n1. -\n2. *\n")
	fmt.Scanln(&operation)

	result := sum(numb1, numb2, operation)
	fmt.Println(result)

}

func sum(a, b, i int) int {
	if i == 0 {
		return a + b
	} else if i == 1 {
		return a - b
	} else if i == 2 {
		return a * b
	} else {
		return 0
	}

}

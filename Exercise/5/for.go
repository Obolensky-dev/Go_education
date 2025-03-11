package main

import "fmt"

func main() {

	var numb1 int

	fmt.Printf("Эта программа вычисляет сумму чисел от 1 до N.\nВведите число N: ")
	fmt.Scanln(&numb1)

	result := sum(numb1)
	fmt.Println(result)

}

func sum(n int) int {
	var result int = 0
	for i := 1; i <= n; i++ {
		result = result + i
	}
	return result
}

package main

import (
	"fmt"
)

// Функция принимает указатель на массив и изменяет его элементы
func modifyArray(arr *[5]int) {
	for i := range arr {
		arr[i] += 10 // Увеличиваем каждый элемент на 10
	}
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println("До изменения:", array)

	modifyArray(&array) // Передаем указатель на массив
	fmt.Println("После изменения:", array)
}

package main

import (
	"fmt"
)

// checkAndPrint проверяет, является ли указатель nil, и безопасно разыменовывает его
func checkAndPrint(ptr *int) {
	if ptr == nil {
		fmt.Println("Ошибка: передан nil-указатель")
		return
	}
	fmt.Println("Значение указателя:", *ptr)
}

func main() {
	var value int = 42
	var ptr1 *int = &value // Указатель на значение
	var ptr2 *int = nil    // nil-указатель

	fmt.Println("Проверка ptr1:")
	checkAndPrint(ptr1)

	fmt.Println("Проверка ptr2:")
	checkAndPrint(ptr2)
}

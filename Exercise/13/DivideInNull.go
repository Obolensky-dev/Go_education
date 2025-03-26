package main

import (
	"fmt"
)

// Функция с возвратом ошибки
func safeDivideWithError(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("ошибка: деление на ноль")
	}
	return a / b, nil
}

// Функция с panic и recover
func safeDivideWithPanic(a, b float64) (result float64) {
	defer func() {
		if err := recover(); err != nil {
			result = 0
			fmt.Println("Обработан критический сбой:", err)
		}
	}()
	if b == 0 {
		panic("критическая ошибка: деление на ноль")
	}
	return a / b
}

func main() {
	// Использование safeDivideWithError
	result, err := safeDivideWithError(10, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Результат деления:", result)
	}
	result, err = safeDivideWithError(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Результат деления:", result)
	}

	// Использование safeDivideWithPanic
	fmt.Println("Результат деления (safeDivideWithPanic):", safeDivideWithPanic(10, 0))
}

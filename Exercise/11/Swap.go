package main

import (
	"fmt"
)

func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func main() {
	a := 5
	b := 12
	fmt.Printf("a = %d, b = %d\n", a, b) // Вывод: 5, 12
	fmt.Println("Magic swap!")
	Swap(&a, &b)
	fmt.Printf("a = %d, b = %d\n", a, b) // Вывод: 12, 5
}

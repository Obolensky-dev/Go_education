package main

import (
	"fmt"
)

func comparePointers(a, b *int) {
	if a == b {
		fmt.Println("Указатели указывают на один и тот же объект")
	} else {
		fmt.Println("Указатели указывают на разные объекты")
	}
}

func main() {
	x, y := 42, 42
	ptr1 := &x
	ptr2 := &x // Указатель на тот же объект
	ptr3 := &y // Указатель на другой объект с тем же значением

	comparePointers(ptr1, ptr2) // Ожидаем: указывают на один объект
	comparePointers(ptr1, ptr3) // Ожидаем: указывают на разные объекты
}

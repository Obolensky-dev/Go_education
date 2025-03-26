package main

import (
	"fmt"
)

func DoubleNumber(x *int) {
	*x *= 2
}

func main() {
	num := 5
	DoubleNumber(&num)
	fmt.Println(num) // Вывод: 10
}

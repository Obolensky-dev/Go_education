/*
Поиск дубликатов: Мы используем карту (map) для отслеживания уже встреченных элементов. Если элемент уже есть в карте, он добавляется в список дубликатов.
Карта: {1: true, 2: true, 3: true, 4: true, 5: true}.
Если seen[n] равно true (уже есть в карте), то в карту не записывается, добавляем n в срез дубликатов: duplicates = append(duplicates, n), если равно seen[n] false, то пишем в карту seen[num] = true
*/

package main

import "fmt"

func findDuplicates(s []int) []int {
	seen := make(map[int]bool)
	duplicates := []int{}

	for _, num := range s {
		if seen[num] {
			duplicates = append(duplicates, num)
		} else {
			seen[num] = true
		}
	}
	return duplicates
}

func main() {
	slice := []int{1, 2, 3, 2, 4, 5, 3}
	duplicates := findDuplicates(slice)
	fmt.Println(duplicates) // [2 3]
}

package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var nums2 []int = make([]int, len(nums))
	lennums := len(nums) - 1
	for _, value := range nums {
		nums2[lennums] = value
		lennums = lennums - 1
	}
	fmt.Println("Срез в обратную сторону:", nums2)

}

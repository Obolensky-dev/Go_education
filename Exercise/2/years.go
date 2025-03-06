package main

import "fmt"

func main() {

	var year int = 2025
	var pi float64 = 3.14
	var name string = "Boris"
	var isGoCool bool = true
	test_var := 1.222
	const Pi = 3.1415
	var stop string

	fmt.Println("Name: ", name)
	fmt.Printf("Year: %d\n", year)
	fmt.Println("Pi: ", pi)
	fmt.Println("Is Go cool?: ", isGoCool)
	fmt.Println("Const pi: ", Pi)
	fmt.Println("Test var: ", test_var)
	fmt.Println("Нажмите Enter для выхода...")
	fmt.Scanln(&stop)

}

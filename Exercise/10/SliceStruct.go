package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	// Создаем слайс структур Person
	people := []Person{
		{Name: "Борис", Age: 29},
		{Name: "Роман", Age: 20},
		{Name: "Кирилл", Age: 27},
	}

	fmt.Printf("Вывод слайса:\n")
	// Итерируемся по слайсу и выводим информацию о каждом человеке
	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}

	fmt.Printf("Добавляем новую структуру в слайс\n")
	// Добавляем новую структуру в слайс
	people = append(people, Person{Name: "Давид", Age: 40})

	fmt.Printf("Вывод слайса:\n")
	// Итерируемся по слайсу и выводим информацию о каждом человеке
	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}

	fmt.Printf("Изменение одного элемента слайса\n")
	//Изменение элемента слайса
	people[1].Age = 26

	fmt.Printf("Вывод слайса:\n")
	// Итерируемся по слайсу и выводим информацию о каждом человеке
	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}

	fmt.Printf("Удаление двух элементов из слайса\n")

	//Удаление элементов из слайса
	people = append(people[:1], people[3:]...)

	fmt.Printf("Вывод слайса:\n")
	// Итерируемся по слайсу и выводим информацию о каждом человеке
	for _, person := range people {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}

}

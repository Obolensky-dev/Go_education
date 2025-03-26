package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	City string
}

func (p *Person) MoveTo(newCity string) {
	p.City = newCity // Изменяем поле структуры через указатель
}

func main() {
	p := Person{"Dave", 40, "London"}
	fmt.Printf("%s is %d years old. He lives in %s\n", p.Name, p.Age, p.City)
	p.MoveTo("Moscow")
	fmt.Printf("%s moved to %s\n", p.Name, p.City)
}

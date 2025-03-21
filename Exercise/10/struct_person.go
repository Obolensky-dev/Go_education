package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	City string
}

func (p Person) Introduce() {
	fmt.Printf("Привет, меня зовут %s, мне %d лет, я живу в %s.\n", p.Name, p.Age, p.City)
}

func main() {
	p := Person{"Борис", 29, "Москва"}
	p.Introduce()
}

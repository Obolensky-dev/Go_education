package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	City string
}

func main() {
	p1 := Person{"Борис", 29, "Москва"}
	p2 := Person{"Роман", 20, "Казань"}
	fmt.Println(AreEqual(p1, p2)) // false
	p1 = Person{"Борис", 29, "Москва"}
	p2 = Person{"Борис", 29, "Москва"}
	fmt.Println(AreEqual(p1, p2)) // true
}

func AreEqual(p1, p2 Person) bool {
	return p1 == p2
}

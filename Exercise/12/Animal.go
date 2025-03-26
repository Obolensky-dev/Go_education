package main

import "fmt"

type Animal interface {
	Sound()
	Move()
}

func AnimalInfo(animal Animal) {
	animal.Sound()
	animal.Move()
}

type Dog struct{}
type Cat struct{}
type Bird struct{}

func (d Dog) Move() {
	fmt.Println("Ходит")
}
func (c Cat) Move() {
	fmt.Println("Ходит")
}
func (b Bird) Move() {
	fmt.Println("Летает")
}
func (d Dog) Sound() {
	fmt.Println("Лает")
}
func (c Cat) Sound() {
	fmt.Println("Мяукает")
}
func (b Bird) Sound() {
	fmt.Println("Щебечет")
}

func main() {

	Dog := Dog{}
	Cat := Cat{}
	Bird := Bird{}
	AnimalInfo(Dog)
	AnimalInfo(Cat)
	AnimalInfo(Bird)
}

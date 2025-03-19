package main

import (
	"fmt"
)

// Создание карты с ключами типа string и значениями типа string
var words = map[string]string{
	"apple": "яблоко",
	"cat":   "кошка",
	"dog":   "собака",
}

func main() {

	var funcnumb int
	var stop string

	fmt.Println("Это приложение переводчик с английского на русский (он немного маловат и знает только 3 слова:")
	fmt.Println("apple, cat, dog")
	fmt.Println("но ты можешь добавить свои слова в словарь, правда их сохранение для дальнейшего использования не реализовано, так что для проверки их не закрывай приложение")
	fmt.Println("Выбери число с тем функционалом который ты хочешь выполнить:")

	funcnumb = -1
	for funcnumb != 0 {
		fmt.Printf("\n0 - Выход\n1 - Переводчик\n2 - Добавить слово\n")
		fmt.Scanln(&funcnumb)

		switch funcnumb {
		case 0:
			break
		case 1:
			treanslate()
		case 2:
			dictionaryadd()
		default:
			fmt.Println("Не выбрано")
		}
	}
	fmt.Println("Нажмите Enter для выхода...")
	fmt.Scanln(&stop)
}

func treanslate() {
	fmt.Println("Введи слово на английском для перевода, либо введи 0 для выхода в меню:")
	for word := ""; word != "0"; {
		fmt.Scanln(&word)

		if word == "0" {
			break
		}
		value, exists := words[word]
		if exists {
			fmt.Println(value)
		} else {
			fmt.Println("Такого слова нет в переводчике")
		}
	}
}

func dictionaryadd() {
	var engword string
	var ruword string
	fmt.Println("Введите слово на английском для добавления в словать, либо 0 для выхода в меню:")
	for engword != "0" {
		fmt.Scanln(&engword)
		if engword == "0" {
			break
		}
		fmt.Println("Введите его перевод на русский:")
		fmt.Scanln(&ruword)
		words[engword] = ruword
	}

}

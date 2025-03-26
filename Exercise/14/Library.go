package main

import (
	"fmt"
)

// Структура для книги
type Book struct {
	Title   string
	Author  string
	IsTaken bool
}

// Структура для пользователя
type User struct {
	Name  string
	Books []*Book // Список книг, взятых пользователем
}

// Интерфейс управления библиотекой
type LibraryManager interface {
	AddBook(title, author string)
	BorrowBook(user *User, title string) error
	ReturnBook(user *User, title string) error
	ListBooks()
}

// Структура библиотеки
type Library struct {
	Books []*Book
}

// Добавление книги в библиотеку
func (l *Library) AddBook(title, author string) {
	book := &Book{Title: title, Author: author, IsTaken: false}
	l.Books = append(l.Books, book)
}

// Выдача книги пользователю
func (l *Library) BorrowBook(user *User, title string) error {
	for _, book := range l.Books {
		if book.Title == title && !book.IsTaken {
			book.IsTaken = true
			user.Books = append(user.Books, book)
			fmt.Printf("%s взял книгу \"%s\"\n", user.Name, title)
			return nil
		}
	}
	return fmt.Errorf("книга \"%s\" недоступна", title)
}

// Возврат книги в библиотеку
func (l *Library) ReturnBook(user *User, title string) error {
	for i, book := range user.Books {
		if book.Title == title {
			book.IsTaken = false
			user.Books = append(user.Books[:i], user.Books[i+1:]...) // Удаляем из списка пользователя
			fmt.Printf("%s вернул книгу \"%s\"\n", user.Name, title)
			return nil
		}
	}
	return fmt.Errorf("у %s нет книги \"%s\"", user.Name, title)
}

// Вывод списка книг в библиотеке
func (l *Library) ListBooks() {
	fmt.Println("Список книг в библиотеке:")
	for _, book := range l.Books {
		status := "доступна"
		if book.IsTaken {
			status = "занята"
		}
		fmt.Printf("- \"%s\" by %s [%s]\n", book.Title, book.Author, status)
	}
}

func main() {
	library := &Library{}
	library.AddBook("1984", "Джордж Оруэлл")
	library.AddBook("Мастер и Маргарита", "Михаил Булгаков")
	library.AddBook("Преступление и наказание", "Федор Достоевский")

	user := &User{Name: "Алексей"}

	library.ListBooks()

	// Читатель берет книгу
	library.BorrowBook(user, "1984")
	library.BorrowBook(user, "Мастер и Маргарита")

	// Проверяем список книг после выдачи
	library.ListBooks()

	// Читатель возвращает книгу
	library.ReturnBook(user, "1984")

	// Итоговый список книг
	library.ListBooks()
}

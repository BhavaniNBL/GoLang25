package main

import (
	"fmt"
)

type Book struct {
	Title           string
	Author          string
	Pages           int
	CopiesAvailable int
}

func CreateBook(title, author string, pages, copies int) *Book {
	return &Book{
		Title:           title,
		Author:          author,
		Pages:           pages,
		CopiesAvailable: copies,
	}
}

func (b Book) Display() {
	fmt.Printf("Title: %s\n", b.Title)
	fmt.Printf("Author: %s\n", b.Author)
	fmt.Printf("Pages: %d\n", b.Pages)
	fmt.Printf("CopiesAvailable: %d\n", b.CopiesAvailable)
}

func (b *Book) Borrow() string {

	if b.CopiesAvailable > 0 {
		b.CopiesAvailable--
		fmt.Println("Copies after borrowing", *b)
		return "borrow successful"
	} else {
		return "borrow unsuccessful"
	}

}

func (b *Book) ReturnBook() {
	if b.CopiesAvailable != 0 {
		b.CopiesAvailable++
		fmt.Println("Copies after returning", *b)
	}
}

func SwapTitles(b1, b2 *Book) {
	temp := b1.Title
	b1.Title = b2.Title
	b2.Title = temp

	fmt.Println("Swapped titles \n", "Book1:", b1.Title, " ", "Book2:", b2.Title)
}

func main() {

	b1 := CreateBook("The Thousand splendid suns", "hossain", 275, 1)

	b2 := CreateBook("Freedom from the Known", "J.K", 185, 9)

	fmt.Println("Book1 Details:")
	b1.Display()
	fmt.Println("Book2 Details:")
	b2.Display()
	fmt.Println("---------------")

	brwFromBook1 := b1.Borrow()

	fmt.Println("Book1", brwFromBook1)
	fmt.Println("---------------------------------------------")
	brwFromBook2 := b2.Borrow()
	fmt.Println("Book2", brwFromBook2)
	fmt.Println("---------------------------------------------")

	b1.ReturnBook()
	b2.ReturnBook()

	SwapTitles(b1, b2)

}

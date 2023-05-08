package data

import (
	"fmt"
	"sync"
)

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{1, "The Lord of the Rings: The Fellowship of the Ring", true},
	{2, "The Lord of the Rings: The Two Towers", false},
	{3, "The Lord of the Rings: The Return of the King", false},
	{4, "The Hobbit", true},
	{5, "The Silmarillion", false},
	{6, "The Children of Húrin", false},
	{7, "Unfinished Tales", false},
	{8, "The History of Middle-earth", false},
	{9, "Bilbo's Last Song", false},
	{10, "The Fall of Arthur", false},
}

func findBook(id int, m *sync.Mutex) (int, *Book) {
	index := -1
	var book *Book

	m.Lock()
	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
		}
	}
	m.Unlock()

	return index, book
}

func FinishBook(id int, m *sync.Mutex) {
	i, book := findBook(id, m)

	if i < 0 {
		return
	}

	m.Lock()
	book.Finished = true
	books[i] = book
	m.Unlock()

	fmt.Printf("Book %s finished\n", book.Title)
}

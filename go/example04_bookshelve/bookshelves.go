package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Bookshelf struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

// Book describes a book on a bookworm's shelf.
type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Id     string `json:"id"`
}

func (b Book) String() string {
	return fmt.Sprintf("\n\nBook name: %s\nBook author: %s\n\n", b.Title, b.Author)
}

func load(file string) ([]Bookshelf, error) {
	var bookshelves []Bookshelf

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Decode the file and store the content in the variable bookworms.
	err = json.NewDecoder(f).Decode(&bookshelves)
	if err != nil {
		return nil, err
	}

	return bookshelves, nil
}

func findCommonBooks(bookshelves []Bookshelf) []Book {
	common := map[Book]int{}

	for _, bookshelf := range bookshelves {
		for _, book := range bookshelf.Books {
			common[book]++
		}
	}

	books := []Book{}
	for book, no := range common {
		if no == len(bookshelves) {
			books = append(books, book)
		}
	}
	return books
}

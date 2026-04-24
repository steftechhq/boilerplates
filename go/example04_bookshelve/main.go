package main

import (
	"fmt"
	"os"
)

func loadBookworms(filePath string) ([]Bookshelf, error) {
	return load(filePath)
}

func main() {

	bookworms, err := loadBookworms("testdata/book_collection.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	fmt.Println(bookworms)

	fmt.Println(findCommonBooks(bookworms))

}

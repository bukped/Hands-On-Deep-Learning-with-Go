package main

import (
	"fmt"
)

// create main function to execute the program
func main() {
	type book struct {
		IDBook int
		Name   string
	}

	// Membuat instance struct
	p := book{IDBook: 171, Name: "Buku Pemrograman Go"}

	// Mengubah nilai properti struct
	p.IDBook = 171

	// Akses nilai properti struct
	id := p.IDBook

	// Conth embedding struct
	type Sales struct {
		book
		StorePosition string
	}

	Store := Sales{
		book:          book{IDBook: 171, Name: "Buku Pemrograman Go"},
		StorePosition: "Store A",
	}

	// Mengubah nilai properti embedded struct
	Store.IDBook = 117

	fmt.Println(id)

}

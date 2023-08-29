package main

import "fmt"

type Book struct {
	Id    uint
	Name  string
	Stock int64
}

func main() {
	var a Book
	a.Id = 112313
	a.Name = " Bahasa Pemrograman Go "
	a.Stock = 30

	fmt.Println("ID buku", a.Id)
	fmt.Println("Nama buku", a.Name)
	fmt.Println("Jumlah buku", a.Stock)

}

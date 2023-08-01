package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func main() {
	// Contoh JSON
	jsonData := `{"id":"11A23","name":"Buku Pemrograman Go","category":"telnologi"}`

	// Unmarshal JSON ke struct
	var book Book
	err := json.Unmarshal([]byte(jsonData), &book)
	if err != nil {
		log.Fatal(err)
	}

	// Mengakses data dalam struct
	fmt.Println("Book ID : ", book.ID)
	fmt.Println("Book Name : ", book.Name)
	fmt.Println("Category : ", book.Category)

	// Marshal struct ke JSON
	newJSON, err := json.Marshal(book)
	if err != nil {
		log.Fatal(err)
	}

	// Cetak hasil JSON
	fmt.Println("JSON", string(newJSON))

}

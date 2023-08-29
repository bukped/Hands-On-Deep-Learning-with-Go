package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Book struct {
	XMLName  xml.Name `xml:"book"`
	ID       string   `xml:"id"`
	Name     string   `xml:"name"`
	Category string   `xml:"category"`
}

func main() {
	// Contoh XML
	xmlData := `<book><id>T11A23</id><name>BukuPemrogramanGo</name><category>Teknologi</category></book>`

	// Unmarshal XML ke struct
	var book Book
	err := xml.Unmarshal([]byte(xmlData), &book)
	if err != nil {
		log.Fatal(err)
	}

	// Mengakses data dalam struct
	fmt.Println("Book ID : ", book.ID)
	fmt.Println("Book Name : ", book.Name)
	fmt.Println("Category : ", book.Category)

	// Marshal struct ke XML
	newXML, err := xml.MarshalIndent(book, "", " ")

	if err != nil {
		log.Fatal(err)
	}

	// Cetak hasil XML
	fmt.Println("XML : ")
	fmt.Println(string(newXML))

}

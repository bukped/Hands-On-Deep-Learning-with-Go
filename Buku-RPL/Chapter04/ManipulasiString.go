package main

import (
	"fmt"
	"strings"
)

// create main function to execute the program
func main() {
	string := "Buku Pemrograman Go"

	// Mengubah string menjadi lowercase atau uppercase
	lower := strings.ToLower(string)
	upper := strings.ToUpper(string)

	// Memisahkan string menjadi slice berdasarkan sepaarator
	split := strings.Split(string, ",")

	// Menggabungkan elemen slice menjadi string dengan separator
	joined := strings.Join(split, "-")

	// Menggantikan substring dalam string
	replaced := strings.Replace(string, "Pemrograman", "Bahasa", -1)

	fmt.Println(lower)
	fmt.Println(upper)
	fmt.Println(joined)
	fmt.Println(replaced)

}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Konversi dari int ke string
	numInt := 42
	strInt := strconv.Itoa(numInt)
	fmt.Println("Hasil konversi int ke string: %s\n", strInt)

	// Konversi dari float64 ke string
	numFloat := 3.14
	strFloat := strconv.FormatFloat(numFloat, 'f', 2, 64)
	fmt.Printf("Hasil konversi float64 ke string: %s\n", strFloat)
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Konversi dari boolean true ke string
	boolTrue := true
	strTrue := strconv.FormatBool(boolTrue)
	fmt.Printf("Hasil konversi boolean %t ke string: \"%s\"\n", boolTrue, strTrue)

	// Konversi dari boolean false ke string
	boolFalse := false
	strFalse := strconv.FormatBool(boolFalse)
	fmt.Printf("Hasil konversi boolean %t ke string: \"%s\"\n", boolFalse, strFalse)

}

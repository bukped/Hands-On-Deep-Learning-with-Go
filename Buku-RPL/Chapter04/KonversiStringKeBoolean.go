package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Konversi dari string "true" ke boolean
	strTrue := "true"
	boolTrue, err := strconv.ParseBool(strTrue)
	if err != nil {
		fmt.Println("Terjadi kesalahan dalam konversi string ke boolean:", err)
		return
	}
	fmt.Printf("Hasil konversi string \"%s\" ke boolean: %t\n", strTrue, boolTrue)

	// Konversi dari string " false" ke boolean
	strFalse := "false"
	boolFalse, err := strconv.ParseBool(strFalse)
	if err != nil {
		fmt.Println("Terjadi kesalahan dalam konversi string ke boolean:", err)
		return
	}
	fmt.Printf("Hasil konversi string \"%s\" ke boolean: %t\n", strFalse, boolFalse)

}

package main

import (
	"fmt"
)

// create main function to execute the program
func main() {
	array1 := [4]int{1, 2, 3, 4}

	// Mengubah nilai elemen pada indeks tertentu
	array1[3] = 12

	// Mengakses nilai elemen pada indeks tertentu
	val := array1[3]

	// Menggabungkan dua array
	array2 := [3]int{6, 7, 8}
	combined := append(array1[:], array2[:]...)

	// Memotong array
	sliced := array1[1:4]

	fmt.Println(val)
	fmt.Println(combined)
	fmt.Println(sliced)

}

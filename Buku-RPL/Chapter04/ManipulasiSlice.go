package main

import "fmt"

// create main function to execute the program
func main() {
	slice := []int{1, 2, 3, 4}
	// Mengubah nilai elemen pada indeks tertentu
	slice[2] = 12

	// Menambahkan elemen ke slice
	slice = append(slice, 5)

	// Menghapus elemen dari slice
	slice = append(slice[:2], slice[3:]...)

	// Memotong slice
	sliced := slice[1:4]

	fmt.Println(sliced)

}

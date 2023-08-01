package main

import "fmt"

// fungsi linearSearch menerima array dan elemen yang ingin dicari
// Mengembalikan indeks elemen jika ditemukan, atau -1 jika tidak ditemukan

func linearSearch(arr []int, target int) int {
	for i, num := range arr {
		if num == target {
			return i // Elemen ditemukan, kembalikan indeks
		}
	}
	return -1 //Elemen tidak ditemukan
}

func main() {
	// Contoh array dan elemen yang ingin dicari
	array := []int{2, 14, 6, 21, 2, 11, 10, 3}
	target := 11

	// Panggil fungsi linearSearch untuk mencari elemen
	result := linearSearch(array, target)

	if result != -1 {
		fmt.Printf("Elemen %d ditemukan di indeks %d\n", target, result)
	} else {
		fmt.Print("Elemen %d tidak ditemukan dalam array\n", target)
	}
}

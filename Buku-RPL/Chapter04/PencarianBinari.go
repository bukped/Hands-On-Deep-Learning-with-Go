package main

import "fmt"

// fungsi binarySearch menerima array terurut dan elemen yang ingin dicari
// Mengembalikan indeks elemen jika ditemukan, atau -1 jika tidak ditemukan

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2
		midVal := arr[mid]

		if midVal == target {
			return mid // Elemen ditemukan, kembalikan indeks
		} else if midVal < target {
			low = mid + 1 // Cari di bagian kanan array
		} else {
			high = mid - 1 // Cari di bagian kiri array
		}
	}

	return -1 // elemen tidak ditemukan
}

func main() {
	// Contoh array terurut dan elemen yang ingin dicari
	array := []int{4, 8, 13, 18, 21, 25, 29}
	target := 25

	// Panggil fungsi binarySearch untuk mencari elemen
	result := binarySearch(array, target)

	if result != -1 {
		fmt.Printf("Elemen %d ditemukan di indeks %d\n", target, result)
	} else {
		fmt.Printf("Elemen %d tidak ditemukan dalam array\n", target)
	}
}

package main

import "fmt"

// fungsi bubblesort menerima array yang ingin diurutkan secara ascending
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Tukar elemen jika elemen saat ini lebih besar dari elemen berikutnya
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// Jika tidak ada pertukaran pada iterasi ini, array sudah terurut
		if !swapped {
			break
		}
	}
}

func main() {
	// Contoh array yang akan diurutkan
	array := []int{19, 7, 22, 78, 45, 17, 98}

	fmt.Println("Array sebelum diurutkan : ", array)
	bubbleSort(array)
	fmt.Println("Array setelah diurutkan : ", array)
}

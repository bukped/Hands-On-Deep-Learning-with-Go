package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1]
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	sortedLeft := quickSort(left)
	sortedRight := quickSort(right)

	return append(append(sortedLeft, pivot), sortedRight...)

}

func main() {
	// Contoh array yang akan diurutkan
	array := []int{63, 7, 22, 78, 45, 17, 98}

	fmt.Println("Array sebelum diurutkan : ", array)
	sorteedArray := quickSort(array)
	fmt.Println("Array setelah diurutkan : ", sorteedArray)
}

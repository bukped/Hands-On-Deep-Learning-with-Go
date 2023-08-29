package main

import (
	"fmt"
)

func main() {

	var arr = [4]int{10, 20, 30, 40}
	target1 := 20
	for i := 0; i < len(arr); i++ {
		if arr[i] == target1 {
			fmt.Println("Element found at index", i)
			break
		}
	}

}

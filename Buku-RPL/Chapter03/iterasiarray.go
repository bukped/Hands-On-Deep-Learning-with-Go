package main

import (
	"fmt"
)

func main() {

	array1 := [4]int{1, 2, 3, 4}
	for i := 0; i < len(array1); i++ {
		fmt.Println(array1[i])
	}

}

package main

import "fmt"

// create main function to execute the program
func main() {
	x := make(map[string]int)

	// Menambahkan key-value ke map
	x["green"] = 6
	x["yellow"] = 4

	// Mengubah nilai value pada key tertentu
	x["green"] = 10

	// Menghapus key dari map
	delete(x, "yellow")

	// Mengakses value pada key tertentu
	val, ok := x["green"]

	// Iterasi melalui map
	for key, value := range x {
		fmt.Println(key, value)
	}

	fmt.Println(val)
	fmt.Println(ok)

}

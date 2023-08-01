package main

import (
	"fmt"
	"os"
)

// create main function to execute the program
func main() {

	//create file
	newfile, err := os.Create("myfile.txt")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	defer newfile.Close()

	text := "Buku Pemrograman Go"
	_, err = newfile.WriteString(text)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}
}

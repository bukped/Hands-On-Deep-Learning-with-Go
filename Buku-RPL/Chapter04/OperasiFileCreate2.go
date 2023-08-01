package main

import (
	"fmt"
	"io/ioutil"
)

// create main function to execute the program
func main() {

	// write the string to the file
	err := ioutil.WriteFile("myfile.txt", []byte("Buku Pemrograman Go"), 0644)
	if err != nil {
		fmt.Println("Failed to write to file : ", err) //print the failed message
		return
	}
	fmt.Println("Wrote to the file 'myfile.txt'.") //print the succes message
}

package main

import (
	"fmt"
	"io/ioutil"
)

// create main function to execute the program
func main() {

	// readfile
	content, err := ioutil.ReadFile("myfile.txt")
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	fmt.Println(string(content)) //print the succes message
}

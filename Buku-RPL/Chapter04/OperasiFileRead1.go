package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	//read file
	newfile, err := os.Open("myfile.txt")
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	defer newfile.Close()

	scanner := bufio.NewScanner(newfile)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error :", err)
	}
}

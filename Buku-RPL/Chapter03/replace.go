package main

import (
	"fmt"
	"strings"
)

func main() {

	z := "Bahasa Pemrograman Go"
	newZ := strings.Replace(z, "Go", "Golang", -1)
	fmt.Println(newZ)

}

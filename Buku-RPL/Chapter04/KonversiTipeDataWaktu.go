package main

import (
	"fmt"
	"time"
)

func main() {
	// Contoh konversi dari time.Time ke string dengan zona waktu
	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05 MST")
	fmt.Println("Waktu dalam bentuk string dengan zona waktu:", timeString)

	// Contoh konversi dari string ke time.Time dengan zona waktu
	timeString2 := "2023-07-19 12:30:00 +0700 WIB"
	parsedTime, err := time.Parse("2006-01-02 15:04:05 -0700 MST", timeString2)
	if err != nil {
		fmt.Println("Terjadi kesalahan dalam konversi string ke time:", err)
		return
	}
	fmt.Println("Waktu dalam bentuk time.Time dengan zona waktu:", parsedTime)
}

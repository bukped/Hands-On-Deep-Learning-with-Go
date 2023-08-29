package main

import (
	"fmt"
	"strings"
)

func main() {

	//Penggabungan String menggunakan operator +
	string1 := "Buku"
	string2 := "Pemrograman Go"
	result1 := string1 + " " + string2
	fmt.Println("Contoh hasil penggabungan string : ", result1)

	//Pemotongan String menggunakan str[:]
	string3 := "Buku Pemrograman Go"
	substring := string3[0:17]
	fmt.Println("Contoh hasil pemotongan string : ", substring)

	//Pencarian String menggunakan contains()
	strContains := strings.Contains(string3, "Pemrograman")
	fmt.Println("Contoh hasil pencarian string : ", strContains)

	//Penggantian String menggunakan replace()
	stringReplace := strings.Replace(string3, "Pemrograman", "Bahasa", -1)
	fmt.Println("Contoh hasil penggantian string : ", stringReplace)

	//Pengubahan string ke huruf besar/huruf kecil menggunakan ToUpper() dan ToLower()
	stringUpper := strings.ToUpper(string3)
	stringLower := strings.ToLower(string3)
	fmt.Println("Contoh hasil pengubahan string ke huruf besar : ", stringUpper)
	fmt.Println("Contoh hasil pengubahan string ke huruf kecil : ", stringLower)

	//Pemisahan String menggunakan split()
	string4 := "Buku-Pemrograman-Go"
	stringSplit := strings.Split(string4, "-")
	fmt.Println("Contoh hasil pemisahan string : ", stringSplit)

	//Penggabungan string dengan separator menggunakan join
	string5 := []string{"Buku", "Pemrograman", "Go"}
	newString := strings.Join(string5, ",")
	fmt.Println("Contoh hasil penggabungan string dengan join : ", newString)

	//Menghilangkan spasi di awal dan di akhir ,enggunakan TrimSpace()
	string6 := "  Buku Pemrograman Go  "
	trimString := strings.TrimSpace(string6)
	fmt.Println("Contoh hasil trim : ", trimString)
}

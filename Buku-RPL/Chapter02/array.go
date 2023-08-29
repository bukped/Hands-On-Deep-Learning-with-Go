package main

import "fmt"

func main() {
	// Contoh deklarasi dan inisialisasi array dengan tipe string
	var buah [5]string
	buah[0] = "apple"
	buah[1] = "grape"
	buah[2] = "banana"
	buah[3] = "guava"
	buah[4] = "avocado"

	// Mengakses dan mencetak nilai array
	fmt.Println("Nilai array jenis buah:", buah)

	// Contoh deklarasi dan inisialisasi array secara langsung
	angka := [3]int{10, 20, 30}

	// Mengakses dan mencetak nilai array
	fmt.Println("Nilai angka:", angka)

	// Mendeklarasikan array tanpa menyebutkan jumlah elemen
	huruf := [...]string{"a", "b", "c", "d", "e"}

	// Mengakses dan mencetak nilai array
	fmt.Println("Huruf-huruf:", huruf)

	// Menghitung panjang array
	panjangbuah := len(angka)
	fmt.Println("Panjang array buah:", panjangbuah)

	// Mengakses elemen array dengan indeks
	fmt.Println("Elemen array angka pada indeks ke-2:", angka[2])

	// Mengubah nilai elemen array dengan indeks
	angka[2] = 15
	fmt.Println("Setelah diubah, elemen array angka pada indeks ke-2:", angka[2])
}

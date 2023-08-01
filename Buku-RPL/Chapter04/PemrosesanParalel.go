package main

import (
	"fmt"
	"sync"
	"time"
)

func process(id int, wg *sync.WaitGroup, ch chan<- string) {
	defer wg.Done()

	// Melakukan pemrosesan tugas
	time.Sleep(2 * time.Second)

	// Mengirim hasil pemrosesan ke channel
	ch <- fmt.Sprintf("Pemrosesan selesai - ID: %d", id)
}

func main() {
	// Membuat channel untuk menerima hadil pemrosesan
	resultCh := make(chan string)

	// Membuat WaitGroup untuk menunggu pemrosesan selesai
	var wg sync.WaitGroup

	// Menetukan jumlah goroutine yang akan dijalankan
	numWorkers := 5

	// Menambahkan jumlah goroutine ke WaitGroup
	wg.Add(numWorkers)

	// Memulai pemrosesan paralel
	for i := 1; i <= numWorkers; i++ {
		go process(i, &wg, resultCh)
	}

	// Menunggu semua goroutine selesai
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Menerima hasil pemrosesan dari channel
	for result := range resultCh {
		fmt.Println(result)
	}

	fmt.Println("Semua Pemrosesan Selesai")

}

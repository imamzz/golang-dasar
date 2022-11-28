package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) { //dijalankan dulu
	defer logging() //memanggil func logging ketika selesai menjalankan func //menggunakan defer diatas/dibagian atas kode
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result", result)
}

func main() {
	runApplication(10)
}

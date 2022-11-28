package main

import "fmt"

func endApp() {
	message := recover() //recover
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR") // ketika panic, kode dibawah dalam func tidak akan dieksekusi
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	// runApp(false) //bisa dicoba
	runApp(true)
	fmt.Println("Imam")
}

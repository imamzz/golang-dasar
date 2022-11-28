package main

import "fmt"

func main() {
	var name = "Imam"

	if name == "Imam" { // akan dieksekusi jika bernilai true aja
		fmt.Println("Hello Imam")
	} else if name == "Zaini" { // akan dieksekusi jika if bernilai false dan else if bernilai true
		fmt.Println("Hi, Zaini")
	} else { // akan dieksekusi jika "if" atau "else if" tidak bernilai true
		fmt.Println("Hi, kamu siapa?")
	}

	// var length = len(name) (pake cara manual)
	if length := len(name); length > 5 { //lebih efisien
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}

package main

import "fmt"

func Ups(i int) interface{} { //interface kosong, adalah kontrak yang didalamnya tidak ada isi kontraknya
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	var data interface{} = Ups(3) //contoh penggunaan
	fmt.Println(data)
}

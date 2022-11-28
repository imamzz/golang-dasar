package main

import "fmt"

func random() interface{} {
	return "Eko"
}

// tipe data harus sama, gaboleh boolean ke string
func main() {
	var result interface{} = random() // interface kosong
	//var resultString string = result.(string) //konverisi dari interface kosong menjadi string
	//fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}

package main

import "fmt"

// func getHello(name string) string {
// 	return "Hello " + name
// kode dibawah return tidak dieksekusi
// }

func getHello(name string) string {
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Eko")
	fmt.Println(result)

	fmt.Println(getHello(""))
}

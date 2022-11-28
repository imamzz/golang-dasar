package main

import "fmt"

func main() {

	name := "Imamza"

	switch name {
	case "Zaini":
		fmt.Println("Hello Zaini")
		fmt.Println("Apa Kabar")
	case "Joko":
		fmt.Println("Hello Joko")
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Kenalan Boleh")
		fmt.Println("Aku Siapa?")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama sudah benar")
	//}

	//switch tanpa
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}

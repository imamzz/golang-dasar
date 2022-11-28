package main

import "fmt"

func main() {

	// var person map[string]string = map[string]string (versi manual)
	person := map[string]string{ //string pertama adalah tipe key kemudian type value
		"name":    "Imam",
		"address": "Ponorogo",
	}

	person["title"] = "Back-end"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//menghapus map
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Imam"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups") //map dan key

	fmt.Println(book)
	fmt.Println(len(book))
}

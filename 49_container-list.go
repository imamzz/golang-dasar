package main

import (
	"container/list"
	"fmt"
)

// bisa google untuk coba list
func main() {
	data := list.New()

	data.PushBack("Imam")
	data.PushBack("Khairul")
	data.PushBack("Zaini")
	data.PushFront("Budi")

	// data.Front().Next() //untuk mengambil data kedua

	// dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// dari belakang ke depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}

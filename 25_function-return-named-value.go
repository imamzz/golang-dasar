package main

import "fmt"

func getFullName2() (firstName string, middleName string, lastName string) {
	firstName = "Imam"
	middleName = "Khairul"
	lastName = "Zaini"

	return
}

func main() {
	a, b, c := getFullName2() //variabel ga harus firstname, middlename, lastname. bisa diganti a b c
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

package main

import "fmt"

func getFullName() (string, string, string) {
	return "Imam", "Khairul", "Zaini"
}

func main() {
	// firstName, middleName, lastname := getFullName()
	firstName, _, _ := getFullName() //cuma mengambil value first name, yg lain diabaikan
	fmt.Println(firstName)
	// fmt.Println(middleName)
	// fmt.Println(lastName)
}

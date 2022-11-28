package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

// Membuat kontrak agar interface dapat digunakan / kaya turunan ga si
func (person Person) GetName() string {
	return person.Name
}interface-kosong.go

type Animal struct {
	Name string
}

// Membuat kontrak
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var eko Person
	eko.Name = "Eko"

	SayHello(eko)

	cat := Animal{
		Name: "Push",
	}
	SayHello(cat)
}

package helper

import "fmt"

var version = 1 //gabisa diakses dari luar package
var Application = "Belajar Golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}

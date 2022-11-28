package main

import "fmt"

type Filter func(string) string

// func sayHelloWithFilter(name string, filter func(string) string) { //cara manual, tidak efisien jika terdapat banyak func
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Eko", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}

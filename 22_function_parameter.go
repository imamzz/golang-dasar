package main

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	// sayHello("Imam", "Khairul") (cara 1)
	firstName := "Imam" // cara 2
	sayHelloTo(firstName, "Khairul")
	sayHelloTo("Budi", "Nugraha")
}

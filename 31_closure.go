package main

import "fmt"

func main() {
	name := "Imam"
	counter := 0

	increment := func() {
		// name = "Budi" // akan merubah scope(name) diatas
		name := "Budi"
		fmt.Println("Increment")
		counter++
		fmt.Println(name)
	}

	increment()
	increment()

	fmt.Println(counter)
	fmt.Println(name)
}

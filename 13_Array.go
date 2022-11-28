package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Imam"
	names[1] = "Khairul"
	names[2] = "Zaini"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{ //array langsung
		90, 91, 92,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names)) //menghitung panjang array
	fmt.Println(len(values))
}

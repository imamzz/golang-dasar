package main

import "fmt"

func main() {
	// cara manual
	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke", counter)
	// 	counter++
	// }

	//lebih simple (post statement)
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"Imam", "Rudi", "Raffi", "Budi", "Joko"}

	//manual
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	person := make(map[string]string)
	person["name"] = "Imam"
	person["title"] = "Back-end"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}

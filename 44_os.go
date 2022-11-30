package main

import (
	"fmt"
	"os"
)

// coba buka go os di google
func main() {
	args := os.Args // penggunaan (go run 44_os.go Imam Khairul)
	fmt.Println("Argument : ")
	fmt.Println(args)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}

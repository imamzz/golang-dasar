package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

//bersungsi jika kita membutuhkan parameter function lain
func main() {

	sayGoodBye := getGoodBye

	result := sayGoodBye("Imam")
	fmt.Println(result)
	fmt.Println(getGoodBye("Imam")) // sama aja kalo pake cara ini

}

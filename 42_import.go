package main

import (
	"fmt"
	"golang-dasar/helper"
)

func main() {
	helper.SayHello("Imam")
	// helper.sayGoodbye("Eko") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
}

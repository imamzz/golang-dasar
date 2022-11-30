package main

import (
	"fmt"
	"strconv"
)

// bisa google strconv untuk dicoba
func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10) // 10 bisa diganti ke binari (2), oktal (8)
	fmt.Println(value)

	valueInt, _ := strconv.Atoi("2000000") //cara gampang
	fmt.Println(valueInt)
}

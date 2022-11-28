package main

import "fmt"

func main() {
	type NoKTP string //alias untuk string, agar tidak menulis string banyak"
	type Married bool

	var noKtpImam NoKTP = "35022784589265"
	var MarriedStatus Married = false
	fmt.Println(noKtpImam)
	fmt.Println(MarriedStatus)
}

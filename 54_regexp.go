package main

import (
	"fmt"
	"regexp"
)

//saran harus dipelajari

// coba di google atau github
// cek video codepolitan
func main() {
	var regex *regexp.Regexp = regexp.MustCompile("i([a-z])m") //reguler expression //depan harus i, belakang harus m, dan harus huruf kecil

	fmt.Println(regex.MatchString("imam"))
	fmt.Println(regex.MatchString("igum"))
	fmt.Println(regex.MatchString("iPam"))

	search := regex.FindAllString("imam ibam ifam iyum iwom", -1)
	fmt.Println(search)
}

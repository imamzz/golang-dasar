package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Imam Khairul", "Imam")) //mencari kata imam
	fmt.Println(strings.Contains("Imam Khairul", "Budi"))

	fmt.Println(strings.Split("Imam Khairul Zaini", " ")) //memotong

	fmt.Println(strings.ToLower("Imam Khairul Zaini")) //mengubah ke lowercase
	fmt.Println(strings.ToUpper("Imam Khairul Zaini"))
	fmt.Println(strings.ToTitle("Imam Khairul Zaini")) //mengubah huruf awalan menjadi besar

	fmt.Println(strings.Trim("      Imam Khairul     ", " "))         //menghapus karakter di kiri dan kanan
	fmt.Println(strings.ReplaceAll("Imam Joko Imam", "Imam", "Budi")) //mengubah kata imam menjadi budi
}

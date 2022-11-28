package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) { 
	address.Country = "Indonesia" //func untuk mengubah negara
}
z
func main() {
	//
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1

	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	// var address4 *Address = &Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1 // & adalah pointer
	var address3 *Address = &address1

	address2.City = "Bandung" // merubah value city

	// address2 mengubah semua value dari Address
	// pointer * memaksa mengubah value address 1 karena mengikuti perubahan pada address2
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"} 
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address) //membuat Address baru, tapi kosong
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address{
		City:     "Subang",
		Province: "Jawa Barat",
		Country:  "",
	}
	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamat)
}

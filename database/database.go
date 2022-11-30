package database

import "fmt"

var connection string

func init() { //dapat langsung dieksekusi meski belum digunakan
	fmt.Println("Init dipanggil")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}

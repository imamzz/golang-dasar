package main

import "fmt"

// Template data
type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) { //struct func/method
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHuuu() {
	fmt.Println("Huuuuuu from", a.Name)
}

func main() {
	var imam Customer
	imam.Name = "Imam"
	imam.Address = "Indonesia"
	imam.Age = 22

	imam.sayHi("Joko")
	imam.sayHuuu()

	//fmt.Println(imam.Name)
	//fmt.Println(imam.Address)
	//fmt.Println(imam.Age)

	// bisa pake yg ini untuk membuat struct
	//joko := Customer{
	//	Name:    "Joko",
	//	Address: "Cirebon",
	//	Age:     35,
	//}
	//fmt.Println(joko)

	// tidak direkomendasikan, karena rawan error
	//budi := Customer{"Budi", "Jakarta", 35}
	//fmt.Println(budi)
}

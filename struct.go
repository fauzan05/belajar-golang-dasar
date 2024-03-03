package main

import "fmt"

type Customer struct {
	Name,Address    string
	Age     int
}

func (customer Customer) getCustomer(bio string){
	fmt.Println("Hello, this customer name is ", customer.Name, bio)
}


func main() {
	var identity Customer
	identity.Name = "Fauzan"
	identity.Address = "Kebumen"

	fmt.Println(identity)
	gadget := Customer{
		Name: "Budi",
		Address: "Kebumen",
		Age: 32,
	}
	fmt.Println(gadget)

	// struct with method
	fauzan := Customer{Name: "Fauzan"}
	fauzan.getCustomer("Makan adalah yang utama")
}

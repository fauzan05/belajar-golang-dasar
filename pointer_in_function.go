package main

import "fmt"

type Address struct {
	Kabupaten string
	Provinsi  string
	Negara    string
}

type Car struct {
	Model string
}

func (car *Car) Owned() {
	car.Model = "You has Owned a " + car.Model
}

func ChangeAddressToIndonesia(address *Address) {
	address.Negara = "Indonesia"
}

func main() {
	NewAddress := Address{"Cilacap", "Jawa Tengah", ""}
	// atau
	var Address Address = Address{"Kebumen", "Jawa Tengah", ""}
	ChangeAddressToIndonesia(&Address)
	ChangeAddressToIndonesia(&NewAddress)
	// ChangeAddressToIndonesia(&Address);
	fmt.Println(Address)
	fmt.Println(NewAddress)

	y := 5

	x := y

	ptr := &x

	fmt.Println("Nilai dari variabel x : ", *ptr)
	fmt.Println("Alamat memori dari variabel x : ", ptr)
	fmt.Println("Alamat memori dari variabel y : ", &y)

	ferrari := Car{"ferrari"}
	ferrari.Owned()
	fmt.Println(ferrari)

}

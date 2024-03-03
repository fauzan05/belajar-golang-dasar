package main

import "fmt"

func hello(name string) (string, string){
	sayhello := "Hello, " + name
	hey := " hey"
	return sayhello, hey
}

func count(a,b int) int{
	return a + b
}

func getFullName()(string, string){
	return "Fauzan", "Nurhidayat"
}

// named return values
func getCompleteName() (firstname, middlename, lastname string) {
	firstname = "fauzan"
	middlename = "nur"
	lastname = "hidayat"
	return lastname, firstname, middlename // urutan return value akan mempengaruhi hasil
}

func main(){
	fmt.Println(hello("Fauzan"))
	fmt.Println(count(1,2))
	firstName, _  := getFullName() // jika tidak ingin menggunakan return value ke 2 gunakan _
	fmt.Println(firstName)
	fmt.Println("/n")
	fmt.Println(getCompleteName())
}
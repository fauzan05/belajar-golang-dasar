package main

import "fmt"

type hasName interface {
	getName() string
	countAge() int
}
func sayHello(value hasName) {
	fmt.Println("Hello", value.getName())
}

type Person struct {
	Name string
	Age int
}

func (person Person) getName() string {
	return person.Name
}

func (person Person) countAge() int {
	return person.Age
}


func main(){
	person := Person{Name: "fauzan", Age: 23}
	sayHello(person)
	fmt.Println(person.getName())
	fmt.Println(person.countAge())

}
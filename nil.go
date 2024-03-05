package main

import "fmt"

// return nil hanya bisa digunakan untuk return value tipe data interface, map, function/closure, slice, pointer, dan channel
func NewMap(name string) map[string]string{
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main(){
	fmt.Println(NewMap("fauzan"))
	kosongan := func (name string) []string {
		if name == "" {
			return nil
		} else {
			return []string{name}
		}
	}
	fmt.Println(kosongan(""))
	check := NewMap("")
	if check == nil {
		fmt.Println("Data kosong")
	} else {
		fmt.Println("Data ada")
	}
}
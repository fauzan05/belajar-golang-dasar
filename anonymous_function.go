package main

import "fmt"

type Blacklist func(string)bool

func registerUser(name string, blacklist Blacklist){
	if blacklist(name) {
		fmt.Println("Anda tidak mendapatkan akses")
	} else {
		fmt.Println("Anda mendapatkan akses")
	}
}

func main(){
	blacklistUser := func(name string) bool{
		return name == "susi"
	}
	
	login := blacklistUser
	registerUser("Anjing", login)
	registerUser("susi", func (name string) bool {
		return name == "susi"
	})
}
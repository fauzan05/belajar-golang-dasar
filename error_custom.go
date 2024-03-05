package main

import (
	"fmt"
	"reflect"
)

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (validate *validationError) Error() string{
	return validate.Message
}

func (notfound *notFoundError) Error() string {
	return notfound.Message
}

func SaveData(id any) error {
	if id == 0 {
		return &validationError{Message: "id not blank"}
	}
	if reflect.TypeOf(id).Kind() != reflect.Int {
		return &notFoundError{Message: "id bukan integer"}
	}

	return nil
}

func main(){
	err := SaveData(-1)
	if err != nil {
		if validationError, ok := err.(*validationError); ok {
			fmt.Println("validation error : ", validationError.Message)
		} else if notFoundError, ok := err.(*notFoundError); ok {
			fmt.Println("not found error : ", notFoundError.Message)
		} else {
			fmt.Println("unknown error : ", err.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
package pembagian

import (
	"errors"
)

func Pembagian(a,b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("nilai harus lebih dari 0")
	} else {
		return a / b, nil
	} 
}
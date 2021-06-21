package isbn

import (
	"github.com/asaskevich/govalidator"
)

func IsValidISBN(isbn string) bool {
	validISBN := govalidator.IsISBN(isbn, 10) //use version 10
	if validISBN == true {
		return true
	}
	return false
}

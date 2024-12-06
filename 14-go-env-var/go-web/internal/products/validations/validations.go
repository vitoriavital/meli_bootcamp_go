package validations

import (
	"errors"
	"strings"
	"go-web/internal/products/model"
	"go-web/internal/products/repository"
	"strconv"
)

func ValidateCodeValue(codeValue string) bool {
	allProducts, err := repository.LoadProducts()
	if err != nil {
		return false
	}
	for _, p := range allProducts {
		if p.CodeValue == codeValue {
			return false
		}
	}

	return true
}

func ValidDaysInMonth(m int, y int, d int) error {
	var validDay int
	switch m {
    case 2:
        if (y % 4 == 0 && y % 100 != 0) || (y % 400 == 0) {
            validDay = 29
        } else {
			validDay = 28
        }
    case 4, 6, 9, 11:
        validDay = 30
    default:
        validDay = 31
    }
	if d > 0 && d <= validDay {
		return nil
	}
	return errors.New("Error: Invalid Day")
}

func ValidateExpiration(expiration string) error {
	fields := strings.Split(expiration, "/")
	d, err := strconv.Atoi(fields[0])
	if err != nil {
		return errors.New("Error: Invalid Day")
	}
	m, err := strconv.Atoi(fields[1])
	if err != nil {
		return errors.New("Error: Invalid Month")
	}
	y, err := strconv.Atoi(fields[2])
	if err != nil {
		return errors.New("Error: Invalid Year")
	}
	if m < 1 || m > 12 {
		return errors.New("Error: Invalid Month")
	}
	if len(fields[2]) != 4 || y < 1 {
		return errors.New("Error: Invalid Year")
	}
	err = ValidDaysInMonth(m, y, d)
	if err != nil {
		return errors.New("Error: Invalid Day")
	}
	return nil
}

func ValidateNewProduct(requestBody model.Product) error {
	validCode := ValidateCodeValue(requestBody.CodeValue)
	if validCode == false {
		return errors.New("Error: Invalid code value")
	}
	dateErr := ValidateExpiration(requestBody.Expiration)
	if dateErr != nil{
		return dateErr
	}
	if requestBody.Name == "" {
		return errors.New("Error: Invalid name")
	}
	if requestBody.Quantity == 0 {
		return errors.New("Error: Invalid quantity")
	}
	if requestBody.Price == 0.0 {
		return errors.New("Error: Invalid price")
	}
	return nil
}
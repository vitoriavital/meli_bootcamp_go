package repository

import (
	"encoding/json"
	"go-web/internal/model"
	"os"
)

func LoadProducts() ([]model.Product, error) {
	file, err := os.Open("docs/db/products.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var result []model.Product
    json.NewDecoder(file).Decode(&result)

    return result, nil
}

func SaveProducts(allProducts []model.Product) error {
	file, err := os.OpenFile("docs/db/products.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        return err
    }
    defer file.Close()
    err = json.NewEncoder(file).Encode(allProducts)
	if err != nil {
		return err
	}
    return nil
}
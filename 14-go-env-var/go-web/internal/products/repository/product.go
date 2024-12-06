package repository

import (
	"encoding/json"
	"go-web/internal/products/model"
	"os"
)

func LoadProducts() (map[int]model.Product, error) {
	file, err := os.Open("docs/db/products.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var result []model.Product
    err = json.NewDecoder(file).Decode(&result)
    if err != nil {
        return nil, err
    }
    
    products := make(map[int]model.Product)
    for _, p := range result {
        products[p.Id] = p
    }

    return products, nil
}

func SaveProducts(allProducts map[int]model.Product) error {
	file, err := os.OpenFile("docs/db/products.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
    if err != nil {
        return err
    }
    defer file.Close()

    var products []model.Product
    for _, p := range allProducts {
        products = append(products, p)
    }
    err = json.NewEncoder(file).Encode(products)
	if err != nil {
        return err
	}
    return nil
}
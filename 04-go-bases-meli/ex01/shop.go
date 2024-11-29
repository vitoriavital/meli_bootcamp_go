package main

import "fmt"

var Products []Product = []Product{{ID: 2, Name: "Sugar", Price: 4.25, Description: "Sugar", Category: "Baking"},
									{ID: 6, Name: "Butter", Price: 8.98, Description: "Unsalted butter", Category: "Dairy"},
									{ID: 7, Name: "Eggs", Price: 10.54, Description: "Eggs", Category: "Dairy"},
									{ID: 9, Name: "Milk", Price: 7.26, Description: "Whole milk", Category: "Dairy"}}

type Product struct {
	ID			int
	Name		string
	Price		float64
	Description	string
	Category	string
}

func (p *Product) Save() {
	Products = append(Products, *p)
}

func (p Product) GetAll() {
	fmt.Println("All products in this list:")
	fmt.Printf("%10s, %5s, %20s, %8s\n", "Product", "Price", "Description", "Category")
	for _, product := range Products {
		fmt.Printf("%10s, %5.02f, %20s, %8s\n", product.Name, product.Price, product.Description, product.Category)
	}
}

func GetById(id int) *Product {
	for _, product := range Products {
		if product.ID == id {
			return &product
		}
	}
	return nil
}

func main() {
	chocolate := Product{
		ID: 0,
		Name: "Chocolate", 
		Price: 9.5,
		Description: "Milk Chocolate",
		Category: "Sweets",
	}
	chocolate.Save()
	flour := Product{
		ID: 1,
		Name: "Flour", 
		Price: 5.5,
		Description: "All-purpose flour",
		Category: "Baking",
	}
	flour.Save()
	chocolate.GetAll()
	item := GetById(7)
	if item != nil {
		fmt.Printf("\nItem at id 7: %s,  %.02f\n", item.Name, item.Price)
	}
}
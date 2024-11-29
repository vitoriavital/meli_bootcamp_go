package main

import "fmt"

type Small struct {
	Cost float64
}

type Medium struct {
	Cost float64
}

type Large struct {
	Cost float64
}

type Product interface {
	Price() float64
}

func (product Small) Price() float64 {
	return product.Cost
}

func (product Medium) Price() float64 {
	return 1.06 * product.Cost
}

func (product Large) Price() float64 {
	return 1.06 * product.Cost + 2500.0
}

func factory(productType string, price float64) Product {
	if productType == "Small" {
		return Small{price}
	} else if productType == "Medium" {
		return Medium{price}
	} else {
		return Large{price}
	}
}



func main() {
	var price float64
	var product Product

	product = factory("Small", 10.5)
	price = product.Price()
	fmt.Printf("%7s| %6s | %8s\n", "Type", "Cost", "Price")
	fmt.Printf("%7s| %6.02f | %8.02f\n", "Small", 10.5, price)

	product = factory("Medium", 245.7)
	price = product.Price()
	fmt.Printf("%7s| %6.02f | %8.02f\n", "Medium", 245.7, price)

	product = factory("Large", 679.8)
	price = product.Price()
	fmt.Printf("%7s| %6.02f | %8.02f\n", "Large", 679.8, price)
}
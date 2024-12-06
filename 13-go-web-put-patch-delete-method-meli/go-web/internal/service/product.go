package service

import (
	"errors"
	"go-web/internal/model"
	"go-web/internal/repository"
	"go-web/internal/validations"
)

type ProductService struct {
	FilePath string
}

func NewProductService(filePath string) *ProductService {
	return &ProductService{FilePath: filePath}
}

func (s *ProductService) GetAllProducts() (map[int]model.Product, error) {
	products, err := repository.LoadProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductById(id int) (*model.Product, error) {
	products, err := repository.LoadProducts()
	_, ok := products[id]
	if ok != true {
		return nil, errors.New("Error: This product doesn't exist!")
	}
	if err != nil {
		return nil, err
	}
	var product *model.Product
	for _, p := range products {
		if p.Id == id {
			product = &p
		}
	}
	return product, nil
}

func (s *ProductService) GetProductsByPrice(priceGt float64) ([]model.Product, error) {
	allProducts, err := repository.LoadProducts()
	if err != nil {
		return nil, err
	}
	var products []model.Product
	for _, product := range allProducts {
		if product.Price >= priceGt {
			products = append(products, product)
		}
	}
	return products, nil
}

func (s *ProductService) CreateProduct(product model.Product) (*model.ResponseBodyProduct, error) {
	allProducts, err := repository.LoadProducts()
	if err != nil {
		return nil, err
	}
	product.Id = len(allProducts) + 1
	validProduct := validations.ValidateNewProduct(product)
	if validProduct != nil {
		return nil, err
	}
	allProducts[product.Id] = product
	err = repository.SaveProducts(allProducts)
	if err != nil {
		return nil, err
	}
	res := model.ResponseBodyProduct{
		Message: "New product created!",
		Product: &product,
		Error:   false,
	}
	return &res, nil
}

func (s *ProductService) UpdateProduct(new model.Product, id int) (*model.ResponseBodyProduct, error) {
	products, err := repository.LoadProducts()
	if err != nil {
		return nil, err
	}
	new.Id = id
	products[id] = new
	err = repository.SaveProducts(products)
	if err != nil {
		return nil, err
	}
	res := model.ResponseBodyProduct{
		Message: "Product updated!",
		Product: &new,
		Error:   false,
	}
	return &res, nil
}

func (s *ProductService) DeleteProduct(id int) (*model.ResponseBodyProduct, error) {
	products, err := repository.LoadProducts()
	if err != nil {
		return nil, err
	}
	delete(products, id)
	err = repository.SaveProducts(products)
	if err != nil {
		return nil, err
	}
	res := model.ResponseBodyProduct{
		Message: "Product deleted!",
		Product: nil,
		Error:   false,
	}
	return &res, nil
}

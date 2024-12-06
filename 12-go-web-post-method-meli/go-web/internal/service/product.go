package service

import (
	"go-web/internal/model"
	"go-web/internal/repository"
	"go-web/pkg/validations"
)

type ProductService struct {
	FilePath string
}

func NewProductService(filePath string) *ProductService {
	return &ProductService{FilePath: filePath}
}

func (s *ProductService) GetAllProducts() ([]model.Product, error) {
	products, err := repository.LoadProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductById(id int) (*model.Product, error) {
	products, err := repository.LoadProducts()
	if err != nil {
		return nil, err
	}
	var product model.Product
	for _, p := range products {
		if p.Id == id {
			product = p
		}
	}
	return &product, nil
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
	err = repository.SaveProducts(append(allProducts, product))
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

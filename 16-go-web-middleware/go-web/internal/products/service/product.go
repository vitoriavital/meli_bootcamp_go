package service

import (
	"errors"
	"go-web/internal/products/model"
	"go-web/internal/products/repository"
	"go-web/internal/products/validations"
)

type ProductService struct {
	Repository *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{Repository: repo}
}

func (s *ProductService) GetAllProducts() (map[int]model.Product, error) {
	products, err := s.Repository.LoadProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) GetProductById(id int) (*model.Product, error) {
	product, err := s.Repository.GetProductById(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) GetProductsByPrice(priceGt float64) (map[int]model.Product, error) {
	products, err := s.Repository.GetProductsByPrice(priceGt)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *ProductService) CreateProduct(product model.Product) (*model.Product, error) {
	allProducts, err := s.Repository.GetAllProducts()
	if err != nil {
		return nil, err
	}
	product.Id = len(allProducts) + 1
	validProduct := validations.ValidateNewProduct(product, s.Repository)
	if validProduct != nil {
		return nil, errors.New("invalid product")
	}
	p, err := s.Repository.CreateProduct(product)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *ProductService) UpdateProduct(new model.Product, id int) (*model.Product, error) {
	new.Id = id
	product, err := s.Repository.UpdateProduct(new, id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(id int) (error) {
	err := s.Repository.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}

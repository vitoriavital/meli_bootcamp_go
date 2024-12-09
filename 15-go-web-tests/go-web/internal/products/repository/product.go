package repository

import (
	"encoding/json"
	"errors"
	"go-web/internal/products/model"
	"os"
)

type ProductRepository struct {
    FilePath string
}

func NewProductRepository(filePath string) *ProductRepository {
	return &ProductRepository{FilePath: filePath}
}

func (r *ProductRepository)LoadProducts() (map[int]model.Product, error) {
	file, err := os.Open(r.FilePath)
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

func (r *ProductRepository) SaveProducts(allProducts map[int]model.Product) error {
	file, err := os.OpenFile(r.FilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
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

func (r *ProductRepository) GetAllProducts() (map[int]model.Product, error) {
	products, err := r.LoadProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) GetProductById(id int) (*model.Product, error) {
	products, err := r.LoadProducts()
	_, ok := products[id]
	if ok != true {
		return nil, errors.New("this product doesn't exist!")
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

func (r *ProductRepository) GetProductsByPrice(priceGt float64) (map[int]model.Product, error) {
	allProducts, err := r.LoadProducts()
	if err != nil {
		return nil, err
	}
	products := make(map[int]model.Product)
	for _, product := range allProducts {
		if product.Price >= priceGt {
            products[product.Id] = product
		}
	}
	return products, nil
}

func (r *ProductRepository) CreateProduct(product model.Product) (*model.Product, error) {
	allProducts, err := r.LoadProducts()
	if err != nil {
		return nil, err
	}
	allProducts[product.Id] = product
	err = r.SaveProducts(allProducts)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) UpdateProduct(new model.Product, id int) (*model.Product, error) {
	products, err := r.LoadProducts()
	if err != nil {
		return nil, err
	}
	products[id] = new
	err = r.SaveProducts(products)
	if err != nil {
		return nil, err
	}
	return &new, nil
}

func (r *ProductRepository) DeleteProduct(id int) (error) {
	products, err := r.LoadProducts()
	if err != nil {
		return err
	}
	delete(products, id)
	err = r.SaveProducts(products)
	if err != nil {
		return err
	}
	return nil
}

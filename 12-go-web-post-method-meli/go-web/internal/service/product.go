package service

type ProductService struct {
	FilePath string
}

func NewProductService(filePath string) *ProductService {
	return &ProductService{FilePath: filePath}
}

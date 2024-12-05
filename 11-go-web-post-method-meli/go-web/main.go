package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"github.com/go-chi/chi/v5"
	"errors"
)

type Product struct {
	Id			int		`json:"id"`
	Name		string	`json:"name"`
	Quantity	int		`json:"quantity"`
	CodeValue	string	`json:"code_value"`
	IsPublished	bool	`json:"is_published"`
	Expiration	string	`json:"expiration"`
	Price		float64	`json:"price"`
}

type RequestBodyProduct struct {
	Name		string	`json:"name"`
	Quantity	int		`json:"quantity"`
	CodeValue	string	`json:"code_value"`
	IsPublished	bool	`json:"is_published"`
	Expiration	string	`json:"expiration"`
	Price		float64	`json:"price"`
}

type ProductData struct {
	Id			int		`json:"id"`
	Name		string	`json:"name"`
	Quantity	int		`json:"quantity"`
	CodeValue	string	`json:"code_value"`
	IsPublished	bool	`json:"is_published"`
	Expiration	string	`json:"expiration"`
	Price		float64	`json:"price"`
}

type ResponseBodyProduct struct {
	Message		string			`json:"message"`
	Product	*Product 	`json:"product,omitempty"`
	Error		bool			`json:"error"`
}

func loadProducts() ([]Product, error) {
	file, err := os.Open("/Users/mlvital/Desktop/bootcamp/meli_bootcamp_go/11-go-web-post-method-meli/go-web/products.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var result []Product
    json.NewDecoder(file).Decode(&result)

    return result, nil
}

func saveProducts(allProducts []Product) error {
	file, err := os.OpenFile("/Users/mlvital/Desktop/bootcamp/meli_bootcamp_go/11-go-web-post-method-meli/go-web/products.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
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

func HandlerPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func validateCodeValue(codeValue string) bool {
	allProducts, err := loadProducts()
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

func daysInMonth(m int, y int, d int) error {
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

func validateExpiration(expiration string) error {
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
	err = daysInMonth(m, y, d)
	if err != nil {
		return errors.New("Error: Invalid Day")
	}
	return nil
}

func validateNewProduct(requestBody RequestBodyProduct) error {
	validCode := validateCodeValue(requestBody.CodeValue)
	if validCode == false {
		return errors.New("Error: Invalid code value")
	}
	dateErr := validateExpiration(requestBody.Expiration)
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

func HandlerCreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	var requestBody RequestBodyProduct
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	validProduct := validateNewProduct(requestBody)
	products, pErr := loadProducts()
	if err != nil || validProduct != nil || pErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {
			fmt.Println(err)
		}
		if validProduct != nil {
			fmt.Println(validProduct)
		}
		if pErr != nil {
			fmt.Println(pErr)
		}
		responseBody := ResponseBodyProduct{
			Message:	"Couldn't Create product!",
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	newProduct := Product{
		Id: 			len(products) + 1,
		Name:			requestBody.Name,
		Quantity:		requestBody.Quantity,
		CodeValue:		requestBody.CodeValue,
		IsPublished:	requestBody.IsPublished,
		Expiration:		requestBody.Expiration,
		Price:			requestBody.Price,
	}
	res := ResponseBodyProduct{
		Message: "New product created!",
		Product: &newProduct,
		Error: false,
	}
	err = saveProducts(append(products, newProduct))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		responseBody := ResponseBodyProduct{
			Message:	"Couldn't save product!",
			Product: nil,
			Error: true,
		}
		json.NewEncoder(w).Encode(responseBody)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func HandlerProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	products, err := loadProducts()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	var msg string
	for _, p := range products {
		msg += fmt.Sprint("id: ", p.Id, " - ", p.Name, " - $", p.Price, " - $", p.Quantity, "\n")
	}
	w.Write([]byte(msg))
}

func HandlerProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	id := chi.URLParam(r, "id")
	products,err := loadProducts()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	nId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	p := products[nId - 1]
	var msg string
	msg += fmt.Sprint("id ", id, " - ", p.Name, " - $", p.Price, " - $", p.Quantity, "\n")
	w.Write([]byte(msg))
}

func HandlerProductSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	
	price := r.URL.Query().Get("priceGt")
	if price == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Missing priceGt"))
		return
	}
	priceGt, err := strconv.ParseFloat(price, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	products,err := loadProducts()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	w.WriteHeader(http.StatusOK)
	var allProducts []Product
	for _, product := range products {
		if product.Price >= priceGt {
			allProducts = append(allProducts, product)
		}
	}
	var msg string
	for _, p := range allProducts {
		msg += fmt.Sprint("id: ", p.Id , " - ", p.Name, " - $", p.Price, " - $", p.Quantity, "\n")
	}
	w.Write([]byte(msg))
}

func ProductRoutes(rt chi.Router) {
	rt.Get("/", HandlerProducts)
	rt.Post("/", HandlerCreateProduct)
	rt.Get("/{id}", HandlerProductById)
	rt.Get("/search", HandlerProductSearch)
}

func main() {
    rt := chi.NewRouter()

	rt.Get("/ping", HandlerPing)
	rt.Route("/products", ProductRoutes)

    if err := http.ListenAndServe(":8080", rt); err != nil {
        panic(err)
    }  
}
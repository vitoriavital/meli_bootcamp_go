package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"os"
	"encoding/json"
	"strconv"
)

type Product struct {
	Id			int		`json:"id"`
	Name		string	`json:"name"`
	Quantity	int		`json:"quantity"`
	CodeValue	string	`json:"code-value"`
	IsPublished	bool	`json:"is-published"`
	Expiration	string	`json:"expiration"`
	Price		float64	`json:"price"`
}

func loadProducts() ([]Product, error) {
	file, err := os.Open("/Users/mlvital/Desktop/bootcamp/meli_bootcamp_go/10-go-web-get-method-meli/go-web/products.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var result []Product
    json.NewDecoder(file).Decode(&result)

    return result, nil
}

func HandlerPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func HandlerProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	products, err := loadProducts()
	if err != nil {
		fmt.Println(err)
		return
	}
	var msg string
	for _, p := range products {
		msg += fmt.Sprint("id: ", p.Id, " - ", p.Name, " - $", p.Price, " - $", p.Quantity, "\n")
	}
	w.Write([]byte(msg))
}

func HandlerProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	id := chi.URLParam(r, "id")
	products,err := loadProducts()
	if err != nil {
		fmt.Println(err)
		return
	}
	nId, err := strconv.Atoi(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	p := products[nId]
	var msg string
	msg += fmt.Sprint("id ", id, " - ", p.Name, " - $", p.Price, " - $", p.Quantity, "\n")
	w.Write([]byte(msg))
}

func HandlerProductSearch(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	price := r.URL.Query().Get("priceGt")
	if price == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Missing priceGt"))
		return
	}
	priceGt, err := strconv.ParseFloat(price, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	products,err := loadProducts()
	if err != nil {
		fmt.Println(err)
		return
	}
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
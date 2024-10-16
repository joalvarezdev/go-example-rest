package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joalvarezdev/go-gpt/model"
	"github.com/joalvarezdev/go-gpt/repository"
)


func GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products, err := repository.GetAllProducts()

	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	product, err := repository.GetProductById(params["id"])

	if err != nil {
		http.Error(w, "Error fetching product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product model.Product

	_ = json.NewDecoder(r.Body).Decode(&product)

	product.Id = strconv.Itoa(rand.Intn(1000000))
	err := repository.CreateProduct(product)

	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    var product model.Product
    _ = json.NewDecoder(r.Body).Decode(&product)
    product.Id = params["id"]
    err := repository.UpdateProduct(product)
    if err != nil {
        http.Error(w, "Error updating product", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(product)
}

// Eliminar un producto
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    err := repository.DeleteProduct(params["id"])
    if err != nil {
        http.Error(w, "Error deleting product", http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode("Product deleted")
}
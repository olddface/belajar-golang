package apiProduct

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stock int    `json:"stock"`
}

var product = []Product{
	{
		Id:    1,
		Nama:  "Product 1",
		Harga: 10000,
		Stock: 10,
	},
	{
		Id:    2,
		Nama:  "Product 2",
		Harga: 20000,
		Stock: 20,
	},
	{
		Id:    3,
		Nama:  "Product 3",
		Harga: 30000,
		Stock: 30,
	},
}

func GetAllProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	} else if r.Method == "POST" {
		// baca data dari request
		// dan masukin ke product
		var newProduct Product
		err := json.NewDecoder(r.Body).Decode(&newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "gagal decode request body",
			})
			return
		}
		newProduct.Id = len(product) + 1
		product = append(product, newProduct)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	}
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	for _, p := range product {
		if p.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
}

func UpdateProductById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var updateProduct Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "gagal decode request body",
		})
		return
	}

	for i := range product {
		if product[i].Id == id {
			updateProduct.Id = id
			product[i] = updateProduct

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "product berhasil di update",
			})
			return
		}
	}
}

func DeleteProductById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/product/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	for i := range product {
		if product[i].Id == id {
			product = append(product[:i], product[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "product berhasil di delete",
			})
			return
		}
	}
}

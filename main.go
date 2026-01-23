package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/apiCategory"
	"kasir-api/apiProduct"
	"net/http"
)

func main() {
	// GET by id
	http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			apiProduct.GetProductById(w, r)
		} else if r.Method == "PUT" {
			apiProduct.UpdateProductById(w, r)
		} else if r.Method == "DELETE" {
			apiProduct.DeleteProductById(w, r)
		}

	})
	// routes product /api/product
	http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
		apiProduct.GetAllProduct(w, r)
	})

	http.HandleFunc("/api/category/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			apiCategory.GetCategoryById(w, r)
		} else if r.Method == "PUT" {
			apiCategory.UpdateCategoryById(w, r)
		} else if r.Method == "DELETE" {
			apiCategory.DeleteCategoryById(w, r)
		}

	})
	// routes Category /api/category
	http.HandleFunc("/api/category", func(w http.ResponseWriter, r *http.Request) {
		apiCategory.GetAllCategory(w, r)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "oke api running",
		})
	})

	fmt.Println("server running di port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}

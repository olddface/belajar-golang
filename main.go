package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Produk struct {
	Id    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stock int    `json:"stock"`
}

var produk = []Produk{
	{
		Id:    1,
		Nama:  "Produk 1",
		Harga: 10000,
		Stock: 10,
	},
	{
		Id:    2,
		Nama:  "Produk 2",
		Harga: 20000,
		Stock: 20,
	},
	{
		Id:    3,
		Nama:  "Produk 3",
		Harga: 30000,
		Stock: 30,
	},
}

func getProductById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	for _, p := range produk {
		if p.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
}

func updateProductById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var updateProduct Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "gagal decode request body",
		})
		return
	}

	for i := range produk {
		if produk[i].Id == id {
			updateProduct.Id = id
			produk[i] = updateProduct

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "produk berhasil di update",
			})
			return
		}
	}
}

func deleteProductById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	for i := range produk {
		if produk[i].Id == id {
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "produk berhasil di delete",
			})
			return
		}
	}
}

func main() {
	// GET by id
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProductById(w, r)
		} else if r.Method == "PUT" {
			updateProductById(w, r)
		} else if r.Method == "DELETE" {
			deleteProductById(w, r)
		}

	})
	// routu product /api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			// baca data dari request
			// dan masukin ke produk
			var newProduk Produk
			err := json.NewDecoder(r.Body).Decode(&newProduk)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode(map[string]string{
					"error": "gagal decode request body",
				})
				return
			}
			newProduk.Id = len(produk) + 1
			produk = append(produk, newProduk)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(produk)
		}
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

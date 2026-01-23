package apiCategory

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	Id   int    `json:"id"`
	Nama string `json:"nama"`
}

var category = []Category{
	{
		Id:   1,
		Nama: "Category 1",
	},
	{
		Id:   2,
		Nama: "Category 2",
	},
	{
		Id:   3,
		Nama: "Category 3",
	},
}

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(category)
	} else if r.Method == "POST" {
		// baca data dari request
		// dan masukin ke product
		var newCategory Category
		err := json.NewDecoder(r.Body).Decode(&newCategory)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"error": "gagal decode request body",
			})
			return
		}
		newCategory.Id = len(category) + 1
		category = append(category, newCategory)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(category)
	}
}

func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	for _, p := range category {
		if p.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
}

func UpdateCategoryById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var updateCategory Category
	err = json.NewDecoder(r.Body).Decode(&updateCategory)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "gagal decode request body",
		})
		return
	}

	for i := range category {
		if category[i].Id == id {
			updateCategory.Id = id
			category[i] = updateCategory

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "category berhasil di update",
			})
			return
		}
	}
}

func DeleteCategoryById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	for i := range category {
		if category[i].Id == id {
			category = append(category[:i], category[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "category berhasil di delete",
			})
			return
		}
	}
}

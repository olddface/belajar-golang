package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

// ubah Config
type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {
	// Read env with viper from OS first
	// Because many provider dont support .env file
	// *******************
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		// then fallback to .env file
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}
	// *******************

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	// Setup database
	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running di", addr)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)
	// Setup routes
	http.HandleFunc("/api/product", productHandler.HandleProducts)
	http.HandleFunc("/api/product/", productHandler.HandleProductByID)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	http.HandleFunc("/api/category", categoryHandler.HandleCategories)
	http.HandleFunc("/api/category/", categoryHandler.HandleCategoryByID)
	// Specific route requested: /category/{category_name}/detail
	http.HandleFunc("/api/category_detail/", categoryHandler.HandleCategoryDetail)

	// GET by id
	// http.HandleFunc("/api/product/", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET" {
	// 		apiProduct.GetProductById(w, r)
	// 	} else if r.Method == "PUT" {
	// 		apiProduct.UpdateProductById(w, r)
	// 	} else if r.Method == "DELETE" {
	// 		apiProduct.DeleteProductById(w, r)
	// 	}

	// })
	// // routes product /api/product
	// http.HandleFunc("/api/product", func(w http.ResponseWriter, r *http.Request) {
	// 	apiProduct.GetAllProduct(w, r)
	// })

	// http.HandleFunc("/categories/", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "GET" {
	// 		apiCategory.GetCategoryById(w, r)
	// 	} else if r.Method == "PUT" {
	// 		apiCategory.UpdateCategoryById(w, r)
	// 	} else if r.Method == "DELETE" {
	// 		apiCategory.DeleteCategoryById(w, r)
	// 	}

	// })
	// // routes Category /category
	// http.HandleFunc("/categories", func(w http.ResponseWriter, r *http.Request) {
	// 	apiCategory.GetAllCategory(w, r)
	// })

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "oke api running",
		})
	})

	err = http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("gagal running server", err)
	}
}

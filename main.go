package main

import (
	"fmt"
	"kasir-api/api"
	"kasir-api/config"
	"kasir-api/config/database"
	"kasir-api/handlers/handlers_impl"
	"kasir-api/repositories/repositories_impl"
	"kasir-api/routes"
	"kasir-api/services/services_impl"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		return
	}
	db, err := database.InitDB(cfg.DBConn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	/*
		API Produk
	*/
	// Init Product Handler
	productRepository := repositories_impl.NewProductRepositoryImpl(db)
	productService := services_impl.NewProductServiceImpl(productRepository)
	productHandler := handlers_impl.NewProductHandlerImpl(productService)

	/*
		API Kategori
	*/
	// GET localhost:8080/api/categories
	// POST localhost:8080/api/categories
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			api.NewKategoriApi().GetAllKategori(w, r)
		} else if r.Method == "POST" {
			api.NewKategoriApi().CreateKategori(w, r)
		}
	})

	// GET localhost:8080/api/categories/{id}
	// PUT localhost:8080/api/categories/{id}
	// DELETE localhost:8080/api/categories/{id}
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			api.NewKategoriApi().GetKategoriByID(w, r)
		} else if r.Method == "PUT" {
			api.NewKategoriApi().UpdateKategori(w, r)
		} else if r.Method == "DELETE" {
			api.NewKategoriApi().DeleteKategori(w, r)
		}
	})

	// Setup routes
	routes.NewRouter(productHandler)

	fmt.Println("server running di localhost:8080")
	err = http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}

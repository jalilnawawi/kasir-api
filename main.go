package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/api"
	"net/http"
)

func main() {
	/*
		API Produk
	*/

	// GET localhost:8080/api/produk/{id}
	// PUT localhost:8080/api/produk/{id}
	// DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			api.NewProdukApi().GetProdukByID(w, r)
		} else if r.Method == "PUT" {
			api.NewProdukApi().UpdateProduk(w, r)
		} else if r.Method == "DELETE" {
			api.NewProdukApi().DeleteProduk(w, r)
		}
	})

	// GET localhost:8080/api/produk
	// POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			api.NewProdukApi().GetAllProduk(w, r)
		} else if r.Method == "POST" {
			api.NewProdukApi().CreateProduk(w, r)
		}
	})

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
			api.NewKategoriApi().GetKategoriById(w, r)
		} else if r.Method == "PUT" {
			api.NewKategoriApi().UpdateKategori(w, r)
		} else if r.Method == "DELETE" {
			api.NewKategoriApi().DeleteKategori(w, r)
		}
	})

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API running",
		})
	})
	fmt.Println("server running di localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}

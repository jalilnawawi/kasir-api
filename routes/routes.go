package routes

import (
	"encoding/json"
	"kasir-api/handlers"
	"net/http"
	"strings"
)

func NewRouter(
	productHandler handlers.ProductHandler,
	categoryHandler handlers.CategoryHandler,
	transactionHandler handlers.TransactionHandler,
	reportHandler handlers.ReportHandler,
) {
	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API running",
		})
	})

	// Product routes
	// GET localhost:8080/api/products
	// POST localhost:8080/api/products
	http.HandleFunc("/api/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetAllProducts(w, r)
		case http.MethodPost:
			productHandler.CreateProduct(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

	// GET localhost:8080/api/products/{id}
	// PUT localhost:8080/api/products/{id}
	// DELETE localhost:8080/api/products/{id}
	http.HandleFunc("/api/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			productHandler.GetProductByID(w, r)
		case http.MethodPut:
			productHandler.UpdateProduct(w, r)
		case http.MethodDelete:
			productHandler.DeleteProduct(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

	// Category routes
	// GET localhost:8080/api/categories
	// POST localhost:8080/api/categories
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			categoryHandler.GetAllCategories(w, r)
		case http.MethodPost:
			categoryHandler.CreateCategory(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

	// GET localhost:8080/api/categories/{id}
	// PUT localhost:8080/api/categories/{id}
	// DELETE localhost:8080/api/categories/{id}
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		if strings.HasSuffix(path, "/products") {
			if r.Method == http.MethodGet {
				categoryHandler.GetProductListByCategoryID(w, r)
				return
			}
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}

		categoryPath := strings.TrimPrefix(path, "/api/categories/")
		if categoryPath != "" && !strings.Contains(path, "/products") {
			switch r.Method {
			case http.MethodGet:
				categoryHandler.GetCategoryByID(w, r)
			case http.MethodPut:
				categoryHandler.UpdateCategory(w, r)
			case http.MethodDelete:
				categoryHandler.DeleteCategory(w, r)
			default:
				http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			}
			return
		}
	})

	// Checkout routes
	// GET localhost:8080/api/checkout
	// POST localhost:8080/api/checkout
	http.HandleFunc("/api/checkout", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			transactionHandler.HandleCheckout(w, r)
		case http.MethodGet:
			transactionHandler.GetAllTransaction(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

	// Report routes
	// GET localhost:8080/api/report/hari-ini
	http.HandleFunc("/api/report", func(w http.ResponseWriter, r *http.Request) {
		reportHandler.GetReportByDate(w, r)
	})

	// GET localhost:8080/api/report/hari-ini
	http.HandleFunc("/api/report/hari-ini", func(w http.ResponseWriter, r *http.Request) {
		reportHandler.GetReportToday(w, r)
	})
}

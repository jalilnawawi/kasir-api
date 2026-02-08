package main

import (
	"fmt"
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
		log.Fatalf("failed to load config: %v", err)
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
	// Init Category Handler
	categoryRepository := repositories_impl.NewCategoryRepositoryImpl(db)
	categoryService := services_impl.NewCategoryServiceImpl(categoryRepository)
	categoryHandler := handlers_impl.NewCategoryHandlerImpl(categoryService)

	/*
		API Transaction
	*/
	transactionRepository := repositories_impl.NewTransactionRepositoryImpl(db)
	transactionService := services_impl.NewTransactionServiceImpl(transactionRepository)
	transactionHandler := handlers_impl.NewTransactionHandlerImpl(transactionService)

	/*
		API Report
	*/
	reportRepository := repositories_impl.NewReportRepositoryImpl(db)
	reportService := services_impl.NewReportServiceImpl(reportRepository)
	reportHandler := handlers_impl.NewReportHandlerImpl(reportService)

	// Setup routes
	routes.NewRouter(productHandler, categoryHandler, transactionHandler, reportHandler)

	fmt.Println("server running di localhost:8080")
	err = http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}

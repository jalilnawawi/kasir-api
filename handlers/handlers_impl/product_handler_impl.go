package handlers_impl

import (
	"encoding/json"
	"kasir-api/error_constant"
	"kasir-api/handlers"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandlerImpl struct {
	service services.ProductService
}

func NewProductHandlerImpl(service services.ProductService) handlers.ProductHandler {
	return &ProductHandlerImpl{service: service}
}

func (h *ProductHandlerImpl) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	products, err := h.service.GetAll(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(products)
	if err != nil {
		http.Error(w, error_constant.ErrFailedGetProduct.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandlerImpl) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, error_constant.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Create(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		http.Error(w, error_constant.ErrFailedAddProduct.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandlerImpl) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, error_constant.ErrInvalidProductID.Error(), http.StatusBadRequest)
		return
	}

	product, err := h.service.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		http.Error(w, error_constant.ErrFailedGetProduct.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandlerImpl) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, error_constant.ErrInvalidProductID.Error(), http.StatusBadRequest)
		return
	}

	var product models.Product
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, error_constant.ErrInvalidRequestBody.Error(), http.StatusBadRequest)
		return
	}

	product.ID = id
	err = h.service.Update(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(product)
	if err != nil {
		http.Error(w, error_constant.ErrFailedUpdateProduct.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ProductHandlerImpl) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, error_constant.ErrInvalidProductID.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
	if err != nil {
		http.Error(w, error_constant.ErrFailedDeleteProduct.Error(), http.StatusInternalServerError)
		return
	}
}

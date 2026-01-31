package handlers_impl

import (
	"encoding/json"
	"kasir-api/handlers"
	"kasir-api/models"
	"kasir-api/services"
	"net/http"
	"strconv"
	"strings"
)

type CategoryHandlerImpl struct {
	service services.CategoryService
}

func NewCategoryHandlerImpl(service services.CategoryService) handlers.CategoryHandler {
	return &CategoryHandlerImpl{service: service}
}

func (h *CategoryHandlerImpl) GetAllCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		http.Error(w, "gagal mendapatkan kategori", http.StatusInternalServerError)
		return
	}
}

func (h *CategoryHandlerImpl) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.Create(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		http.Error(w, "gagal menambahkan kategori", http.StatusInternalServerError)
		return
	}
}

func (h *CategoryHandlerImpl) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid category ID", http.StatusBadRequest)
		return
	}

	category, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		http.Error(w, "gagal mendapatkan kategori", http.StatusInternalServerError)
		return
	}
}

func (h *CategoryHandlerImpl) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid category ID", http.StatusBadRequest)
		return
	}

	var category models.Category
	err = json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	category.ID = id
	err = h.service.Update(&category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(category)
	if err != nil {
		http.Error(w, "gagal memperbarui kategori", http.StatusInternalServerError)
		return
	}
}

func (h *CategoryHandlerImpl) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid category ID", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(map[string]string{"message": "Category deleted successfully"})
	if err != nil {
		http.Error(w, "gagal menghapus kategori", http.StatusInternalServerError)
		return
	}
}

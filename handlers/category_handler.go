package handlers

import "net/http"

type CategoryHandler interface {
	GetAllCategories(w http.ResponseWriter, r *http.Request)
	CreateCategory(w http.ResponseWriter, r *http.Request)
	GetCategoryByID(w http.ResponseWriter, r *http.Request)
	UpdateCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	GetProductListByCategoryID(w http.ResponseWriter, r *http.Request)
}

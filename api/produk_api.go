package api

import (
	"encoding/json"
	"kasir-api/model"
	"net/http"
	"strconv"
	"strings"
)

type ProdukApi interface {
	GetAllProduk(w http.ResponseWriter, r *http.Request)
	CreateProduk(w http.ResponseWriter, r *http.Request)
	GetProdukByID(w http.ResponseWriter, r *http.Request)
	UpdateProduk(w http.ResponseWriter, r *http.Request)
	DeleteProduk(w http.ResponseWriter, r *http.Request)
}

func NewProdukApi() ProdukApi {
	return &ProdukApiImpl{}
}

type ProdukApiImpl struct{}

func (service *ProdukApiImpl) GetAllProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(model.ListProduk)
	if err != nil {
		http.Error(w, "gagal mendapatkan produk", http.StatusInternalServerError)
		return
	}
}

func (service *ProdukApiImpl) CreateProduk(w http.ResponseWriter, r *http.Request) {
	var produkBaru model.Produk
	err := json.NewDecoder(r.Body).Decode(&produkBaru)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	produkBaru.ID = len(model.ListProduk) + 1
	model.ListProduk = append(model.ListProduk, produkBaru)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(produkBaru)
	if err != nil {
		http.Error(w, "gagal menambahkan produk", http.StatusInternalServerError)
		return
	}
}

func (service *ProdukApiImpl) GetProdukByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for _, p := range model.ListProduk {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func (service *ProdukApiImpl) UpdateProduk(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateProduk model.Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop produk, cari id, ganti sesuai data dari request
	for i := range model.ListProduk {
		if model.ListProduk[i].ID == id {
			updateProduk.ID = id
			model.ListProduk[i] = updateProduk

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func (service *ProdukApiImpl) DeleteProduk(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}
	// loop produk cari ID, dapet index yang mau dihapus
	for i, p := range model.ListProduk {
		if p.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			model.ListProduk = append(model.ListProduk[:i], model.ListProduk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})

			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

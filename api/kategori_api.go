package api

import (
	"encoding/json"
	"kasir-api/model"
	"net/http"
	"strconv"
	"strings"
)

type KategoriApi interface {
	GetAllKategori(w http.ResponseWriter, r *http.Request)
	CreateKategori(w http.ResponseWriter, r *http.Request)
	GetKategoriByID(w http.ResponseWriter, r *http.Request)
	UpdateKategori(w http.ResponseWriter, r *http.Request)
	DeleteKategori(w http.ResponseWriter, r *http.Request)
}

func NewKategoriApi() KategoriApi {
	return &KategoriApiImpl{}
}

type KategoriApiImpl struct{}

func (service KategoriApiImpl) GetAllKategori(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(model.ListKategori)
	if err != nil {
		http.Error(w, "gagal mendapatkan kategori", http.StatusInternalServerError)
		return
	}
}

func (service KategoriApiImpl) CreateKategori(w http.ResponseWriter, r *http.Request) {
	var kategoriBaru model.Kategori
	err := json.NewDecoder(r.Body).Decode(&kategoriBaru)
	if err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	kategoriBaru.ID = len(model.ListKategori) + 1
	model.ListKategori = append(model.ListKategori, kategoriBaru)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(kategoriBaru)
	if err != nil {
		http.Error(w, "gagal menambahkan kategori", http.StatusInternalServerError)
		return
	}
}

func (service KategoriApiImpl) GetKategoriByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	for _, p := range model.ListKategori {
		if p.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Kategori belum ada", http.StatusNotFound)
}

func (service KategoriApiImpl) UpdateKategori(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	var updateKategori model.Kategori
	err = json.NewDecoder(r.Body).Decode(&updateKategori)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range model.ListKategori {
		if model.ListKategori[i].ID == id {
			updateKategori.ID = id
			model.ListKategori[i] = updateKategori

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateKategori)
			return
		}
	}
	http.Error(w, "Kategori belum ada", http.StatusNotFound)
}

func (service KategoriApiImpl) DeleteKategori(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	for i, p := range model.ListKategori {
		if p.ID == id {
			model.ListKategori = append(model.ListKategori[:i], model.ListKategori[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})

			return
		}
	}

	http.Error(w, "Kategori belum ada", http.StatusNotFound)
}

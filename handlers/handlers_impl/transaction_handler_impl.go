package handlers_impl

import (
	"encoding/json"
	"kasir-api/error_constant"
	"kasir-api/handlers"
	"kasir-api/models/dto"
	"kasir-api/services"
	"net/http"
)

type TransactionHandlerImpl struct {
	service services.TransactionService
}

func NewTransactionHandlerImpl(service services.TransactionService) handlers.TransactionHandler {
	return &TransactionHandlerImpl{service: service}
}

func (h *TransactionHandlerImpl) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Checkout(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *TransactionHandlerImpl) Checkout(w http.ResponseWriter, r *http.Request) {
	var req dto.CheckoutRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	transaction, err := h.service.Checkout(req.Items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transaction)
	if err != nil {
		return
	}
}

func (h *TransactionHandlerImpl) GetAllTransaction(w http.ResponseWriter, r *http.Request) {
	transactions, err := h.service.GetAllTransaction()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		http.Error(w, error_constant.ErrFailedGetCategory.Error(), http.StatusInternalServerError)
		return
	}
}

package handlers

import "net/http"

type TransactionHandler interface {
	HandleCheckout(w http.ResponseWriter, r *http.Request)
	Checkout(w http.ResponseWriter, r *http.Request)
	GetAllTransaction(w http.ResponseWriter, r *http.Request)
}

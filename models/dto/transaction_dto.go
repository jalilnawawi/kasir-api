package dto

import "time"

type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"`
}

type CheckoutItem struct {
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type TransactionDto struct {
	ID          int                     `json:"id"`
	TotalAmount int                     `json:"total_amount"`
	Details     []TransactionDetailsDto `json:"details"`
	CreatedAt   time.Time               `json:"created_at"`
}

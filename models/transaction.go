package models

type Transaction struct {
	ID          int                  `json:"id"`
	TotalAmount int                  `json:"total_amount"`
	Details     []TransactionDetails `json:"details"`
}

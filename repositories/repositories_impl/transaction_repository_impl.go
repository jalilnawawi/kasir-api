package repositories_impl

import (
	"database/sql"
	"fmt"
	"kasir-api/models"
	"kasir-api/models/dto"
	"kasir-api/repositories"
	"time"
)

type TransactionRepositoryImpl struct {
	db *sql.DB
}

func NewTransactionRepositoryImpl(db *sql.DB) repositories.TransactionRepository {
	return &TransactionRepositoryImpl{db: db}
}

func (r *TransactionRepositoryImpl) CreateTransaction(items []dto.CheckoutItem) (*models.Transaction, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetails, 0)
	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product with ID %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2 AND stock >= $1", item.Quantity, item.ProductID)
		if err != nil {
			return nil, fmt.Errorf("insufficient stock for product %s. Available: %d, Requested: %d", productName, stock, item.Quantity)
		}
		details = append(details, models.TransactionDetails{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id", totalAmount).Scan(&transactionID)
	if err != nil {
		return nil, err
	}

	for i := range details {
		details[i].TransactionID = transactionID
		_, err = tx.Exec("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4)",
			transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal)
		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

func (r *TransactionRepositoryImpl) GetAllTransaction() (*[]dto.TransactionDto, error) {
	query := `
				SELECT
				    t.id as transaction_id,
				    t.total_amount as total_amount,
				    td.product_id as product_id,
				    p.name as product_name,
				    td.quantity as quantity,
				    td.subtotal as subtotal,
				    t.created_at as checkout_at
				FROM transactions t
				JOIN transaction_details td ON td.transaction_id = t.id
				JOIN products p ON p.id = td.product_id
				ORDER BY t.created_at DESC
             `

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	transactionMap := make(map[int]*dto.TransactionDto)

	for rows.Next() {
		var (
			transactionID int
			totalAmount   int
			productID     int
			productName   string
			quantity      int
			subtotal      int
			createdAt     time.Time
		)

		err := rows.Scan(&transactionID, &totalAmount, &productID, &productName, &quantity, &subtotal, &createdAt)
		if err != nil {
			return nil, err
		}

		if _, exists := transactionMap[transactionID]; !exists {
			transactionMap[transactionID] = &dto.TransactionDto{
				ID:          transactionID,
				TotalAmount: totalAmount,
				CreatedAt:   createdAt,
				Details:     make([]dto.TransactionDetailsDto, 0),
			}
		}

		transactionMap[transactionID].Details = append(
			transactionMap[transactionID].Details,
			dto.TransactionDetailsDto{
				TransactionID: transactionID,
				ProductID:     productID,
				ProductName:   productName,
				Quantity:      quantity,
				Subtotal:      subtotal,
			},
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	result := make([]dto.TransactionDto, 0, len(transactionMap))
	for _, transaction := range transactionMap {
		result = append(result, *transaction)
	}

	return &result, nil
}

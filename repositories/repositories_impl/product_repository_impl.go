package repositories_impl

import (
	"database/sql"
	"errors"
	"fmt"
	"kasir-api/models"
	"kasir-api/models/dto"
	"kasir-api/repositories"
)

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepositoryImpl(db *sql.DB) repositories.ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (repo *ProductRepositoryImpl) GetAll() ([]models.Product, error) {
	query := "SELECT id, name, price, stock, category_id FROM products"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		products = append(products, p)
	}

	return products, nil
}

func (repo *ProductRepositoryImpl) Create(produk *models.Product) error {
	query := "INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) returning id"
	err := repo.db.QueryRow(query, produk.Name, produk.Price, produk.Stock, produk.CategoryID).Scan(&produk.ID)
	return err
}

func (repo *ProductRepositoryImpl) GetByID(id int) (*dto.ProductDetailDto, error) {
	query := `
				SELECT 
				    p.id, 
				    p.name, 
				    p.price, 
				    p.stock, 
				    c.name 
				FROM products p JOIN categories c ON c.id = p.category_id 
				WHERE p.id = $1
			 `

	var p dto.ProductDetailDto
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryName)
	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	return &p, nil
}

func (repo *ProductRepositoryImpl) Update(produk *models.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5"
	result, err := repo.db.Exec(query, produk.Name, produk.Price, produk.Stock, produk.CategoryID, produk.ID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return nil
}

func (repo *ProductRepositoryImpl) Delete(id int) error {
	query := "DELETE FROM products WHERE id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return errors.New("produk tidak ditemukan")
	}

	return err
}

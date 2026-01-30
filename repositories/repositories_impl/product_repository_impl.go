package repositories_impl

import (
	"database/sql"
	"errors"
	"fmt"
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductRepositoryImpl struct {
	db *sql.DB
}

func NewProductRepositoryImpl(db *sql.DB) repositories.ProductRepository {
	return &ProductRepositoryImpl{db: db}
}

func (repo *ProductRepositoryImpl) GetAll() ([]models.Product, error) {
	query := "SELECT id, name, price, stock FROM products"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	products := make([]models.Product, 0)
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		products = append(products, p)
	}

	return products, nil
}

func (repo *ProductRepositoryImpl) Create(produk *models.Product) error {
	query := "INSERT INTO products (name, price, stock) VALUES ($1, $2, $3) returning id"
	err := repo.db.QueryRow(query, produk.Name, produk.Price, produk.Stock).Scan(&produk.ID)
	return err
}

func (repo *ProductRepositoryImpl) GetByID(id int) (*models.Product, error) {
	query := "SELECT id, name, price, stock FROM products WHERE id = $1"

	var p models.Product
	err := repo.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Price, &p.Stock)
	if err == sql.ErrNoRows {
		return nil, errors.New("produk tidak ditemukan")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	return &p, nil
}

func (repo *ProductRepositoryImpl) Update(produk *models.Product) error {
	query := "UPDATE products SET name = $1, price = $2, stock = $3 WHERE id = $4"
	result, err := repo.db.Exec(query, produk.Name, produk.Price, produk.Stock, produk.ID)
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

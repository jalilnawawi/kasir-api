package repositories_impl

import (
	"database/sql"
	"fmt"
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryRepositoryImpl struct {
	db *sql.DB
}

func NewCategoryRepositoryImpl(db *sql.DB) repositories.CategoryRepository {
	return &CategoryRepositoryImpl{db: db}
}

func (repo *CategoryRepositoryImpl) GetAll() ([]models.Category, error) {
	query := "select id, name, description from categories"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer rows.Close()

	categories := make([]models.Category, 0)
	for rows.Next() {
		var cat models.Category
		err := rows.Scan(&cat.ID, &cat.Name, &cat.Description)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		categories = append(categories, cat)
	}

	return categories, nil
}

func (repo *CategoryRepositoryImpl) Create(kategori *models.Category) error {
	query := "insert into categories (name, description) values ($1, $2) returning id"
	err := repo.db.QueryRow(query, kategori.Name, kategori.Description).Scan(&kategori.ID)
	return err
}

func (repo *CategoryRepositoryImpl) GetByID(id int) (*models.Category, error) {
	query := "select id, name, description from categories where id = $1"

	var cat models.Category
	err := repo.db.QueryRow(query, id).Scan(&cat.ID, &cat.Name, &cat.Description)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}

	return &cat, nil
}

func (repo *CategoryRepositoryImpl) Update(kategori *models.Category) error {
	query := "update categories set name = $1, description = $2 where id = $3"
	result, err := repo.db.Exec(query, kategori.Name, kategori.Description, kategori.ID)
	if err != nil {
		return fmt.Errorf("failed to execute update: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("no category found with id %d", kategori.ID)
	}

	return nil
}

func (repo *CategoryRepositoryImpl) Delete(id int) error {
	query := "delete from categories where id = $1"
	result, err := repo.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to execute delete: %v", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to retrieve affected rows: %v", err)
	}

	if rows == 0 {
		return fmt.Errorf("no category found with id %d", id)
	}
	return nil
}

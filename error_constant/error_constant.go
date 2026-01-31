package error_constant

import "errors"

var (
	ErrInvalidRequestBody = errors.New("invalid request body")

	ErrProductNotFound     = errors.New("product not found")
	ErrInvalidProductID    = errors.New("invalid product id")
	ErrFailedGetProduct    = errors.New("failed to get product")
	ErrFailedAddProduct    = errors.New("failed to add product")
	ErrFailedUpdateProduct = errors.New("failed to update product")
	ErrFailedDeleteProduct = errors.New("failed to delete product")

	ErrCategoryNotFound     = errors.New("category not found")
	ErrInvalidCategoryID    = errors.New("invalid category id")
	ErrFailedGetCategory    = errors.New("failed to get category")
	ErrFailedAddCategory    = errors.New("failed to add category")
	ErrFailedUpdateCategory = errors.New("failed to update category")
	ErrFailedDeleteCategory = errors.New("failed to delete category")
)

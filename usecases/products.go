package usecases

import (
	"fmt"
	"tech-test/models"
	"tech-test/repository"
)

// ProductsUsecasesInterface ...
type ProductsUsecasesInterface interface {
	Create(m *models.Product) error
	GetByID(id int, m *models.Product) error
	GetAll(limit, lastID string) ([]models.Product, error)
	UpdateByID(id int, m *models.Product) (rowAffected int, err error)
	DeleteByID(id int) (rowAffected int, err error)
}

// Products Usecases
type Products struct{}

// NewProductsUsecase ...
func NewProductsUsecase() ProductsUsecasesInterface {
	return Products{}
}

// Create ...
func (v Products) Create(m *models.Product) error {
	if m.Title == "" {
		return fmt.Errorf("usecases: title cannot be empty string")
	}
	if m.Description == "" {
		return fmt.Errorf("usecases: description cannot be empty string")
	}
	if m.Image == "" {
		return fmt.Errorf("usecases: image cannot be empty string")
	}
	return repository.NewProductsRepository().Create(m)
}

// GetByID ...
func (v Products) GetByID(id int, m *models.Product) error {
	if id <= 0 {
		return fmt.Errorf("usecases: id cannot be 0 or below")
	}
	return repository.NewProductsRepository().GetByID(id, m)
}

// GetAll ...
func (v Products) GetAll(limit, lastID string) ([]models.Product, error) {
	ml, err := repository.NewProductsRepository().GetAll(limit, lastID)
	if err != nil {
		return nil, fmt.Errorf("usecases: GetAll : %s", err.Error())
	}

	return ml, nil
}

// UpdateByID ...
func (v Products) UpdateByID(id int, m *models.Product) (int, error) {
	if m.Title == "" {
		return -1, fmt.Errorf("usecases: title cannot be empty string")
	}
	if m.Description == "" {
		return -1, fmt.Errorf("usecases: description cannot be empty string")
	}
	if m.Rating == 0 {
		return -1, fmt.Errorf("usecases: rating cannot be 0")
	}
	if m.Image == "" {
		return -1, fmt.Errorf("usecases: image cannot be empty string")
	}
	return repository.NewProductsRepository().UpdateByID(id, m)
}

// DeleteByID ...
func (v Products) DeleteByID(id int) (int, error) {
	if id <= 0 {
		return -1, fmt.Errorf("usecases: id cannot be 0 or below")
	}

	return repository.NewProductsRepository().DeleteByID(id)
}

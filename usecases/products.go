package usecases

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"tech-test/cache"
	"tech-test/models"
	"tech-test/repository"
	"time"
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

	key := fmt.Sprintf("%s:%d", "PRODUCT_GETBYID", id)
	r, err := cache.GetDB().Get(context.Background(), key).Result()
	if err == nil {
		err = json.Unmarshal([]byte(r), &m)
		if err != nil {
			log.Printf("Cannot unmarshal: %s", err.Error())
		}
		return nil
	} else {
		err = repository.NewProductsRepository().GetByID(id, m)
		if err != nil {
			return err
		}
	}

	b, err := json.Marshal(m)
	if err != nil {
		log.Printf("Cannot marshal: %s", err.Error())
	} else {
		_, err = cache.GetDB().Set(context.Background(), key, string(b), time.Minute).Result()
		if err != nil {
			log.Printf("Cannot store to cache: %s", err.Error())
		}
	}

	return nil
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

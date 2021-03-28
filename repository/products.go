package repository

import (
	"database/sql"
	"fmt"
	"tech-test/database"
	"tech-test/models"
	"time"
)

// ProductsRepositoryInterface ...
type ProductsRepositoryInterface interface {
	Create(m *models.Product) error
	GetByID(id int, m *models.Product) error
	GetAll(limit, lastID string) ([]models.Product, error)
	UpdateByID(id int, m *models.Product) (rowAffected int, err error)
	DeleteByID(id int) (rowAffected int, err error)
}

// Products Repository
type Products struct{}

// NewProductsRepository ...
func NewProductsRepository() ProductsRepositoryInterface {
	return Products{}
}

// Create ...
func (v Products) Create(m *models.Product) error {

	query := `
		INSERT INTO products (title, description, rating, image, created_at, updated_at)
		VALUES ($1,$2,$3,$4,$5,$6)
		RETURNING id`
	statement, err := database.GetPostGresDBconn().Prepare(query)

	if err != nil {
		return fmt.Errorf("repository: cannot prepare query to database: %s", err.Error())
	}

	defer statement.Close()

	now := time.Now()
	insertedID := 0
	err = statement.QueryRow(m.Title, m.Description, m.Rating, m.Image, now, now).Scan(&insertedID)
	if err != nil {
		return fmt.Errorf("repository: cannot exec query to database: %s", err.Error())
	}
	m.ID = insertedID // assign last inserted row id to p.ID

	return nil
}

// GetByID ...
func (v Products) GetByID(id int, m *models.Product) error {

	query := `
		SELECT id, title, description, rating, image, created_at, updated_at FROM products
		WHERE id = $1`
	statement, err := database.GetPostGresDBconn().Prepare(query)

	if err != nil {
		return fmt.Errorf("repository: cannot prepare query to database: %s", err.Error())
	}

	defer statement.Close()

	err = statement.QueryRow(id).Scan(
		&m.ID,
		&m.Title,
		&m.Description,
		&m.Rating,
		&m.Image,
		&m.CreatedAt,
		&m.UpdatedAt)
	if err != nil {
		return fmt.Errorf("repository: cannot query row to database: %s", err.Error())
	}

	return nil
}

// GetAll ...
func (v Products) GetAll(limit, lastID string) ([]models.Product, error) {

	query := "SELECT id, title, description, rating, image, created_at, updated_at FROM products WHERE id > $1 ORDER BY rating DESC LIMIT $2"

	statement, err := database.GetPostGresDBconn().Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("repository: cannot prepare query to database: %s", err.Error())
	}
	defer statement.Close()

	rows, err := statement.Query(lastID, limit)
	if err != nil {
		return nil, fmt.Errorf("repository: cannot execute query to database: %s", err.Error())
	}
	defer rows.Close()

	var ml []models.Product
	for rows.Next() {
		var m models.Product
		var updatedAt sql.NullString
		err = rows.Scan(
			&m.ID,
			&m.Title,
			&m.Description,
			&m.Rating,
			&m.Image,
			&m.CreatedAt,
			&m.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("repository: cannot scan rows: %s", err.Error())
		}
		if updatedAt.Valid {
			tu, err := time.Parse(time.RFC3339, updatedAt.String)
			if err != nil {
				return nil, fmt.Errorf("repository: cannot parse update_at time: %s", err.Error())
			}
			m.UpdatedAt = tu
		}
		ml = append(ml, m)
	}

	return ml, nil
}

// UpdateByID ...
func (v Products) UpdateByID(id int, m *models.Product) (int, error) {
	query := `
	UPDATE products
	SET title = $1, 
		description = $2,
		rating = $3,
		image = $4,
		updated_at = $5
	WHERE id = $6`

	statement, err := database.GetPostGresDBconn().Prepare(query)
	if err != nil {
		return -1, fmt.Errorf("repository: cannot prepare query to database: %s", err.Error())
	}
	defer statement.Close()

	r, err := statement.Exec(m.Title, m.Description, m.Rating, m.Image, m.UpdatedAt, id)
	if err != nil {
		return -1, fmt.Errorf("repository: cannot exec query to database: %s", err.Error())
	}

	idInt64, err := r.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("repository: there is no row affected: %s", err.Error())
	}

	return int(idInt64), nil
}

// DeleteByID ...
func (v Products) DeleteByID(id int) (int, error) {
	query := `
	DELETE FROM products 
	WHERE id = $1`
	statement, err := database.GetPostGresDBconn().Prepare(query)
	if err != nil {
		return -1, fmt.Errorf("repository: cannot prepare query to database: %s", err.Error())
	}

	defer statement.Close()

	r, err := statement.Exec(id)
	if err != nil {
		return -1, fmt.Errorf("repository: cannot exec query to database: %s", err.Error())
	}

	idInt64, err := r.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("repository: cannot get row affected: %s", err.Error())
	}

	return int(idInt64), nil
}

package models

import (
	"time"
)

type Product struct {
	ID          int
	Title       string
	Description string
	Rating      int
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

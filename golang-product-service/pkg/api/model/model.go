package model

import (
	"encoding/json"
	"time"
)

type Review struct {
	ID        int       `json:"-" db:"id"`
	Name      string    `json:"name" db:"name" validate:"required"`
	UUID      string    `json:"_id" db:"uuid" validate:"required"`
	Rating    float64   `json:"rating" db:"rating" validate:"required"`
	Comment   string    `json:"comment" db:"comment" validate:"required"`
	UserID    string    `json:"userId" db:"user_id" validate:"required"`
	ProductID int       `json:"-" db:"product_id" validate:"required"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type Product struct {
	ID           int       `json:"-" db:"id"`
	UUID         string    `json:"_id" db:"uuid" validate:"required"`
	Name         string    `json:"name" db:"name" validate:"required"`
	Image        string    `json:"image" db:"image" validate:"required"`
	Brand        string    `json:"brand" db:"brand" validate:"required"`
	Category     string    `json:"category" db:"category" validate:"required"`
	Description  string    `json:"description" db:"description" validate:"required"`
	Reviews      []Review  `json:"reviews" db:"reviews"`
	Rating       float64   `json:"rating" db:"rating" validate:"required"`
	NumReviews   int       `json:"numReviews" db:"num_reviews" validate:"required"`
	Price        float64   `json:"price" db:"price" validate:"required"`
	CountInStock int       `json:"countInStock" db:"count_in_stock" validate:"required"`
	UserID       *string   `json:"userId" db:"user_id" validate:"required"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}

// MarshalJSON initializes nil slices and then marshals the bag to JSON
func (p Product) MarshalJSON() ([]byte, error) {
	type Alias Product

	a := struct {
		Alias
	}{
		Alias: (Alias)(p),
	}

	if a.Reviews == nil {
		a.Reviews = make([]Review, 0)
	}

	return json.Marshal(a)
}

type ProductPagination struct {
	Products []Product `json:"products"`
	Page     int       `json:"page"`
	Pages    int       `json:"pages"`
}

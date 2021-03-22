package repository

import (
	"context"

	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/model"
)

type ReviewRepository interface {
	FindReviewByProductIDAndUserID(ctx context.Context, productID int, userID string) (*model.Review, error)
	FindReviewByProductID(ctx context.Context, productID int) ([]model.Review, error)
}

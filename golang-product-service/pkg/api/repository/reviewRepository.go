package repository

import "github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/model"

type ReviewRepository interface {
	FindReviewByProductIDAndUserID(productID int, userID string) (*model.Review, error)
	FindReviewByProductID(productID int) ([]model.Review, error)
}

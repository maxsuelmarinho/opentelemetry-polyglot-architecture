package repository

import (
	"context"

	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/model"
)

type ProductRepository interface {
	GetProducts(ctx context.Context, keyword string, offset, limit int) ([]model.Product, error)
	GetProductsCount(keyword string) (int, error)
	GetProductByID(ctx context.Context, uuid string) (*model.Product, error)
	CreateProductReview(ctx context.Context, review *model.Review) error
	GetTopProducts(ctx context.Context) ([]model.Product, error)
	UpdateProductReviewDetails(ctx context.Context, product *model.Product) error
}

package repository

import "github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/model"

type ProductRepository interface {
	GetProducts(keyword string, offset, limit int) ([]model.Product, error)
	GetProductsCount(keyword string) (int, error)
	GetProductByID(uuid string) (*model.Product, error)
	CreateProductReview(review *model.Review) error
	GetTopProducts() ([]model.Product, error)
	UpdateProductReviewDetails(product *model.Product) error
}

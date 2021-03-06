package service

import (
	"errors"
	"math"

	apperror "github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/error"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/model"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/repository"
)

type ProductService interface {
	GetProducts(keyword string, page, pageSize int) (*model.ProductPagination, error)
	GetProductByID(uuid string) (*model.Product, error)
	CreateProductReview(uuid string, dto model.CreateReviewDTO) error
	GetTopProducts() ([]model.Product, error)
}

func NewProductService(repository repository.Repository) ProductService {
	return &productService{
		repository: repository,
	}
}

type productService struct {
	repository repository.Repository
}

func (s *productService) GetProducts(keyword string, page, pageSize int) (*model.ProductPagination, error) {
	products, err := s.repository.GetProducts(keyword, pageSize*(page-1), pageSize)
	if err != nil {
		return nil, err
	}

	count, err := s.repository.GetProductsCount(keyword)
	if err != nil {
		return nil, err
	}

	return &model.ProductPagination{
		Products: products,
		Page:     page,
		Pages:    int(math.Ceil(float64(count) / float64(pageSize))),
	}, nil
}

func (s *productService) GetProductByID(uuid string) (*model.Product, error) {
	return s.repository.GetProductByID(uuid)
}

func (s *productService) CreateProductReview(uuid string, dto model.CreateReviewDTO) error {
	product, err := s.repository.GetProductByID(uuid)
	if err != nil {
		return err
	}

	review, err := s.repository.FindReviewByProductIDAndUserID(product.ID, dto.UserID)
	if err != nil && !errors.Is(err, apperror.ErrReviewNotFound) {
		return err
	}

	if review != nil {
		return apperror.ErrProductAlreadyReviewd
	}

	review = &model.Review{
		UserID:    dto.UserID,
		ProductID: product.ID,
		Name:      dto.UserName,
		Rating:    dto.Rating,
		Comment:   dto.Comment,
	}

	if err := s.repository.CreateProductReview(review); err != nil {
		return err
	}

	reviews, err := s.repository.FindReviewByProductID(product.ID)
	if err != nil {
		return err
	}

	product.NumReviews = len(reviews)
	rating := 0.0
	for _, r := range reviews {
		rating += r.Rating
	}
	product.Rating = rating / float64(product.NumReviews)
	if err := s.repository.UpdateProductReviewDetails(product); err != nil {
		return err
	}

	return nil
}

func (s *productService) GetTopProducts() ([]model.Product, error) {
	return s.repository.GetTopProducts()
}

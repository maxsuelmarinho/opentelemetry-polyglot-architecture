package service

import (
	"context"
	"errors"
	"math"

	apperror "github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/error"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/model"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/repository"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type ProductService interface {
	GetProducts(ctx context.Context, keyword string, page, pageSize int) (*model.ProductPagination, error)
	GetProductByID(ctx context.Context, uuid string) (*model.Product, error)
	CreateProductReview(ctx context.Context, uuid string, dto model.CreateReviewDTO) error
	GetTopProducts(ctx context.Context) ([]model.Product, error)
}

func NewProductService(repository repository.Repository) ProductService {
	return &productService{
		repository: repository,
	}
}

type productService struct {
	repository repository.Repository
}

func (s *productService) GetProducts(ctx context.Context, keyword string, page, pageSize int) (*model.ProductPagination, error) {
	products, err := s.repository.GetProducts(ctx, keyword, pageSize*(page-1), pageSize)
	if err != nil {
		return nil, err
	}

	count, err := s.repository.GetProductsCount(ctx, keyword)
	if err != nil {
		return nil, err
	}

	return &model.ProductPagination{
		Products: products,
		Page:     page,
		Pages:    int(math.Ceil(float64(count) / float64(pageSize))),
	}, nil
}

func (s *productService) GetProductByID(ctx context.Context, uuid string) (*model.Product, error) {
	return s.repository.GetProductByID(ctx, uuid)
}

func (s *productService) CreateProductReview(ctx context.Context, uuid string, dto model.CreateReviewDTO) error {
	tracer := otel.Tracer(viper.GetString("APP_NAME"))
	var span trace.Span
	ctx, span = tracer.Start(ctx, "service.CreateProductReview")
	defer span.End()

	span.AddEvent("Find product by id")
	product, err := s.repository.GetProductByID(ctx, uuid)
	if err != nil {
		return err
	}
	span.AddEvent("Product found")
	span.SetAttributes(attribute.String("product.name", product.Name))

	span.AddEvent("Validating whether the user already reviewed this product")
	review, err := s.repository.FindReviewByProductIDAndUserID(ctx, product.ID, dto.UserID)
	if err != nil && !errors.Is(err, apperror.ErrReviewNotFound) {
		return err
	}

	if review != nil {
		err := apperror.ErrProductAlreadyReviewd
		span.RecordError(err)
		return err
	}

	review = &model.Review{
		UserID:    dto.UserID,
		ProductID: product.ID,
		Name:      dto.UserName,
		Rating:    dto.Rating,
		Comment:   dto.Comment,
	}

	span.AddEvent("Creating product review")
	if err := s.repository.CreateProductReview(ctx, review); err != nil {
		return err
	}

	span.AddEvent("Finding product reviews")
	reviews, err := s.repository.FindReviewByProductID(ctx, product.ID)
	if err != nil {
		return err
	}

	span.AddEvent("Updating product rating")
	product.NumReviews = len(reviews)
	rating := 0.0
	for _, r := range reviews {
		rating += r.Rating
	}
	product.Rating = rating / float64(product.NumReviews)
	if err := s.repository.UpdateProductReviewDetails(ctx, product); err != nil {
		return err
	}
	span.SetAttributes(attribute.Float64("product.rating", product.Rating))

	return nil
}

func (s *productService) GetTopProducts(ctx context.Context) ([]model.Product, error) {
	return s.repository.GetTopProducts(ctx)
}

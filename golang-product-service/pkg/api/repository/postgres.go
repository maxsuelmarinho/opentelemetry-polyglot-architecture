package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	apperror "github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/error"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/model"
)

type Repository interface {
	ProductRepository
	ReviewRepository
}

func NewRepository(db *sqlx.DB) Repository {
	return &postgresRepository{
		db: db,
	}
}

type postgresRepository struct {
	db *sqlx.DB
}

func (r *postgresRepository) GetProducts(keyword string, offset, limit int) ([]model.Product, error) {
	var products []model.Product
	selectClause := `SELECT p.* FROM products p`
	pagination := "OFFSET ? LIMIT ?"
	var values []interface{}

	var whereClause string
	if keyword != "" {
		values = append(values, "%"+keyword+"%")
		whereClause = "WHERE p.name ILIKE ?"
	}
	values = append(values, offset)
	values = append(values, limit)

	queryTemplate := fmt.Sprintf("%s %s %s", selectClause, whereClause, pagination)
	query, args, err := sqlx.In(queryTemplate, values...)
	if err != nil {
		return nil, err
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query)
	if err := r.db.Select(&products, query, args...); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	for _, p := range products {
		reviews, err := r.FindReviewByProductID(p.ID)
		if err != nil {
			return nil, err
		}
		p.Reviews = reviews
	}

	return products, nil
}

func (r *postgresRepository) GetProductsCount(keyword string) (int, error) {
	var count int
	selectClause := `SELECT COUNT(1) FROM products p`
	var values []interface{}
	var whereClause string
	if keyword != "" {
		values = append(values, "%"+keyword+"%")
		whereClause = "WHERE p.name ILIKE ?"
	}

	queryTemplate := fmt.Sprintf("%s %s", selectClause, whereClause)
	query, args, err := sqlx.In(queryTemplate, values...)
	if err != nil {
		return 0, err
	}

	query = sqlx.Rebind(sqlx.DOLLAR, query)
	if err := r.db.QueryRow(query, args...).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func (r *postgresRepository) GetProductByID(uuid string) (*model.Product, error) {
	var product model.Product
	query := `SELECT p.* FROM products p WHERE uuid = $1`

	if err := r.db.Get(&product, query, uuid); err != nil {
		if err == sql.ErrNoRows {
			return nil, apperror.ErrProductNotFound
		}
		return nil, err
	}

	reviews, err := r.FindReviewByProductID(product.ID)
	if err != nil {
		return nil, err
	}

	product.Reviews = reviews

	return &product, nil
}

func (r *postgresRepository) CreateProductReview(review *model.Review) error {
	query := `INSERT INTO reviews (name, rating, comment, user_id, product_id)
		VALUES (:name, :rating, :comment, :user_id, :product_id)`
	_, err := r.db.NamedExec(query, &review)
	return err
}

func (r *postgresRepository) GetTopProducts() ([]model.Product, error) {
	var products []model.Product
	query := `SELECT p.* FROM products p ORDER BY p.rating DESC LIMIT 3`
	if err := r.db.Select(&products, query); err != nil {
		return nil, err
	}

	for _, p := range products {
		reviews, err := r.FindReviewByProductID(p.ID)
		if err != nil {
			return nil, err
		}
		p.Reviews = reviews
	}

	return products, nil
}

func (r *postgresRepository) FindReviewByProductIDAndUserID(productID int, userID string) (*model.Review, error) {
	var review model.Review
	query := `
		SELECT r.* FROM reviews r
		WHERE r.product_id = $1 AND r.user_id = $2`

	if err := r.db.Get(&review, query, productID, userID); err != nil {
		if err == sql.ErrNoRows {
			return nil, apperror.ErrReviewNotFound
		}
		return nil, err
	}

	return &review, nil
}

func (r *postgresRepository) FindReviewByProductID(productID int) ([]model.Review, error) {
	var reviews []model.Review
	query := `SELECT r.* FROM reviews r WHERE r.product_id = $1`

	if err := r.db.Select(&reviews, query, productID); err != nil {
		return nil, err
	}

	return reviews, nil
}

func (r *postgresRepository) UpdateProductReviewDetails(product *model.Product) error {
	query := `
		UPDATE products SET rating = $1, num_reviews = $2
		WHERE id = $3`
	if _, err := r.db.Exec(query, product.Rating, product.NumReviews, product.ID); err != nil {
		return err
	}

	return nil
}

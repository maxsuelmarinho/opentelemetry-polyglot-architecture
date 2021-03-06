package error

import "errors"

var (
	ErrProductNotFound       = errors.New("Product not found")
	ErrReviewNotFound        = errors.New("Review not found")
	ErrProductAlreadyReviewd = errors.New("Product already reviewed")
)

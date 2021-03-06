package model

type CreateReviewDTO struct {
	UserID   string  `json:"userId"`
	UserName string  `json:"userName"`
	Rating   float64 `json:"rating"`
	Comment  string  `json:"comment"`
}

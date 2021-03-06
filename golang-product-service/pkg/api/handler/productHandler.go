package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/model"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/service"
	log "github.com/sirupsen/logrus"
)

type ProductHandler interface {
	GetProducts(w http.ResponseWriter, r *http.Request)
	GetProductByID() http.Handler
	CreateProductReview() http.Handler
	GetTopProducts(w http.ResponseWriter, r *http.Request)
}

func NewProductHandler(service service.ProductService, logger *log.Logger) ProductHandler {
	return &productHandler{
		service: service,
		logger:  logger,
	}
}

type productHandler struct {
	service service.ProductService
	logger  *log.Logger
}

func (h *productHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	pageSize := 2
	page := 1

	pageNumber := r.URL.Query().Get("pageNumber")
	if value, err := strconv.Atoi(pageNumber); err == nil {
		page = value
	}

	keyword := r.URL.Query().Get("keyword")

	products, err := h.service.GetProducts(keyword, page, pageSize)
	if err != nil {
		handleError(err, w)
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		handleError(err, w)
		return
	}
}

func (h *productHandler) GetProductByID() http.Handler {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		vars := mux.Vars(r)
		uuid := vars["id"]
		product, err := h.service.GetProductByID(uuid)
		if err != nil {
			return err
		}

		if err := json.NewEncoder(w).Encode(product); err != nil {
			return err
		}
		return nil
	}, h.logger)
}

func (h *productHandler) CreateProductReview() http.Handler {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		var dto model.CreateReviewDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			return err
		}
		defer r.Body.Close()

		vars := mux.Vars(r)
		uuid := vars["id"]

		if err := h.service.CreateProductReview(uuid, dto); err != nil {
			return err
		}

		response := map[string]string{
			"message": "Review added",
		}

		respondWithJSON(w, response, http.StatusCreated)
		return nil
	}, h.logger)
}

func (h *productHandler) GetTopProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.service.GetTopProducts()
	if err != nil {
		handleError(err, w)
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		handleError(err, w)
		return
	}
}

func respondWithJSON(w http.ResponseWriter, payload interface{}, code int) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/model"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/service"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
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
	ctx := r.Context()
	var span trace.Span
	ctx, span = otel.Tracer(viper.GetString("APP_NAME")).Start(ctx, "GetProducts")
	defer span.End()
	endpointKV := httpPathKey.String("/api/v1/products")
	httpRequestsCounter.Add(ctx, 1, endpointKV)

	pageSize := 2
	page := 1

	pageNumber := r.URL.Query().Get("pageNumber")
	if value, err := strconv.Atoi(pageNumber); err == nil {
		page = value
	}

	keyword := r.URL.Query().Get("keyword")

	products, err := h.service.GetProducts(ctx, keyword, page, pageSize)
	if err != nil {
		span.SetStatus(codes.Error, err.Error())
		handleError(err, w)
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		span.SetStatus(codes.Error, err.Error())
		handleError(err, w)
		return
	}

	span.SetStatus(codes.Ok, "Success")
}

func (h *productHandler) GetProductByID() http.Handler {
	return errorHandler(func(w http.ResponseWriter, r *http.Request) error {
		vars := mux.Vars(r)
		uuid := vars["id"]
		ctx := r.Context()
		product, err := h.service.GetProductByID(ctx, uuid)
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
		ctx := r.Context()
		tracer := otel.Tracer(viper.GetString("APP_NAME"))
		var span trace.Span
		ctx, span = tracer.Start(ctx, "CreateProductReview")
		defer span.End()

		var dto model.CreateReviewDTO
		if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
			return err
		}
		defer r.Body.Close()

		vars := mux.Vars(r)
		uuid := vars["id"]

		span.AddEvent("Creating product review", trace.WithAttributes(attribute.String("product.id", uuid)))
		if err := h.service.CreateProductReview(ctx, uuid, dto); err != nil {
			return err
		}

		response := map[string]string{
			"message": "Review added",
		}

		span.AddEvent("Product review created")
		respondWithJSON(w, response, http.StatusCreated)
		return nil
	}, h.logger)
}

func (h *productHandler) GetTopProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	products, err := h.service.GetTopProducts(ctx)
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

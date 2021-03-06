package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/handler"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/repository"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/service"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/logger"
	"github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/util"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func StartServer() {
	logger := logger.CreateLoggerInstance()
	database := repository.NewDatabase(logger)
	database.Initialize()
	repository := repository.NewRepository(database.GetConnection())
	service := service.NewProductService(repository)

	healthCheckHandler := handler.NewHealthCheckHandler()
	productHandler := handler.NewProductHandler(service, logger)

	r := mux.NewRouter()
	r.Methods(http.MethodGet).Path("/health").HandlerFunc(healthCheckHandler.HealthCheck)
	r.Methods(http.MethodGet).Path("/api/products").HandlerFunc(productHandler.GetProducts)
	r.Methods(http.MethodGet).Path("/api/products/top").HandlerFunc(productHandler.GetTopProducts)
	r.Methods(http.MethodGet).Path("/api/products/{id}").Handler(productHandler.GetProductByID())
	r.Methods(http.MethodPost).Path("/api/products/{id}/reviews").Handler(productHandler.CreateProductReview())

	serverPort := viper.GetInt("SERVER_PORT")
	httpServer := &http.Server{Addr: fmt.Sprintf(":%d", serverPort), Handler: r}

	util.HandleSigterm(func() {
		if err := database.GetConnection().Close(); err != nil {
			logger.Error(errors.Wrap(err, "error on closing database connection"))
		} else {
			logger.Debug("database connection has been closed")
		}

		if err := httpServer.Close(); err != nil {
			logger.Error(errors.Wrap(err, "error on closing http server"))
		} else {
			logger.Debug("http server has been closed")
		}
	})

	logger.Info("product service started on port ", serverPort)
	logger.Fatal(httpServer.ListenAndServe())
}

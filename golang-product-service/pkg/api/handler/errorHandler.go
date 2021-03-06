package handler

import (
	"errors"
	"fmt"
	"net/http"

	apperror "github.com/maxsuelmarinho/ecommerce-example/golang-product-service/pkg/api/error"
	log "github.com/sirupsen/logrus"
)

type appError struct {
	Error   error
	Message string
	Code    int
}

type appHandler func(w http.ResponseWriter, r *http.Request) error

func errorHandler(fn appHandler, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			logger.Error(err)
			e := handlerError(err)
			respondWithError(w, e.Message, e.Code)
			return
		}
	})
}

func handlerError(err error) appError {
	if errors.Is(err, apperror.ErrProductNotFound) {
		return appError{err, err.Error(), http.StatusNotFound}
	}

	if errors.Is(err, apperror.ErrProductAlreadyReviewd) {
		return appError{err, err.Error(), http.StatusBadRequest}
	}

	return appError{err, "request process failed", http.StatusInternalServerError}
}

func respondWithError(w http.ResponseWriter, message string, httpCode int) {
	payload := map[string]string{
		"message": message,
	}

	respondWithJSON(w, payload, httpCode)
}

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, err.Error())))
}

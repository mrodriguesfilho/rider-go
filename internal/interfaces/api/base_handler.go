package api

import (
	"encoding/json"
	"net/http"
)

type BaseHandler struct {
}

func NewBaseHandle() *BaseHandler {
	return &BaseHandler{}
}

func (b *BaseHandler) HasDecodeError(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}

	w.WriteHeader(http.StatusBadRequest)
	msg := struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	}

	json.NewEncoder(w).Encode(msg)

	return false
}

func (b *BaseHandler) HasUseCaseError(w http.ResponseWriter, err error) bool {

	if err == nil {
		return false
	}

	w.WriteHeader(http.StatusBadRequest)

	msg := struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	}

	json.NewEncoder(w).Encode(msg)

	return true
}

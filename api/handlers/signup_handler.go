package handlers

import (
	"encoding/json"
	"net/http"
	"rider-go/internal/usecase"
)

type SignUpHandler struct {
	SignUpUseCase *usecase.SignUpUseCase
	*BaseHandler
}

func NewSignUpHandler(signUpUseCase *usecase.SignUpUseCase) *SignUpHandler {
	return &SignUpHandler{
		SignUpUseCase: signUpUseCase,
		BaseHandler:   NewBaseHandle(),
	}
}

func (h *SignUpHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var signUpInput usecase.SignUpInput
	err := json.NewDecoder(r.Body).Decode(&signUpInput)

	if h.BaseHandler.HasDecodeError(w, err) {
		return
	}

	output, err := h.SignUpUseCase.Execute(signUpInput)

	if h.BaseHandler.HasUseCaseError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}

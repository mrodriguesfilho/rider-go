package handlers

import (
	"encoding/json"
	"net/http"
	"rider-go/internal/usecase"
)

type GetAccountHandler struct {
	GetAccountUseCase *usecase.GetAccountUsecase
	*BaseHandler
}

func NewGetAccountHandler(getAccountUseCase *usecase.GetAccountUsecase) *GetAccountHandler {
	return &GetAccountHandler{
		GetAccountUseCase: getAccountUseCase,
		BaseHandler:       NewBaseHandle(),
	}
}

func (g *GetAccountHandler) Handle(w http.ResponseWriter, r *http.Request) {

	email := r.URL.Query().Get("email")

	if email == "" {
		http.Error(w, "email parameter is required", http.StatusBadRequest)
		return
	}

	getAccountInput := usecase.GetAccountInput{
		Email: email,
	}

	account, err := g.GetAccountUseCase.Execute(getAccountInput)

	if g.BaseHandler.HasUseCaseError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(account)
}

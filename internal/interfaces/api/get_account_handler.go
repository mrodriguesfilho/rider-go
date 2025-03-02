package api

import (
	"encoding/json"
	"net/http"
	"rider-go/internal/application/usecase"
)

type GetAccount struct {
	GetAccountUseCase *usecase.GetAccount
	*BaseHandler
}

func NewGetAccountHandler(getAccountUseCase *usecase.GetAccount) *GetAccount {
	return &GetAccount{
		GetAccountUseCase: getAccountUseCase,
		BaseHandler:       NewBaseHandle(),
	}
}

func (g *GetAccount) Handle(w http.ResponseWriter, r *http.Request) {

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

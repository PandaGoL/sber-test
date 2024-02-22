package v1

import (
	"net/http"
	"sber-test/internal/services/deposit"

	"github.com/gorilla/mux"
)

// Router - cтруктура данных для HTTP API v1
type Router struct {
	mainRouter *mux.Router
	router     *mux.Router
	Service    deposit.DepositServicer
}

// InitAPI - функция инициализирует HTTP API версии 1
func NewAPI(mainRouter *mux.Router, s deposit.DepositServicer, middleware func(next http.Handler) http.Handler) {
	sr := &Router{
		mainRouter: mainRouter,
		router:     mainRouter.PathPrefix("/v1").Subrouter(),
		Service:    s,
	}
	//sr.router.Use(middleware)

	sr.router.HandleFunc("/deposit", sr.CalculateDeposit).Methods("POST")
}

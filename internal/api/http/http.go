package http

import (
	"context"
	"net/http"
	"time"

	v1 "sber-test/internal/api/http/v1"
	"sber-test/internal/services/deposit"
	"sber-test/pkg/options"

	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

// Server - основной объект HTTP API сервера
type Server struct {
	router        *mux.Router
	service       deposit.DepositServicer
	httpSrv       *http.Server
	shutdownState bool
}

// Init - функция инициализирует и возвращает HTTP API сервер
func Init(opt *options.Options, depositService deposit.DepositServicer) *Server {
	s := &Server{
		router:  mux.NewRouter(),
		service: depositService,
	}

	//ручка сервиса расчета депозита
	v1.NewAPI(s.router, depositService, s.middlewareRequestID)

	lr := s.router
	s.httpSrv = &http.Server{
		Handler: lr,
		Addr:    opt.APIAddr,
	}

	return s
}

// Serve - функция для запуска HTTP API сервера
func (s *Server) Serve() (err error) {
	err = s.httpSrv.ListenAndServe()
	return
}

// Stop - функция остановки работы сервера
func (s *Server) Stop() (err error) {
	s.setShutdownState(true)
	// ждем завершения или просто "гасим" сервер
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	return s.httpSrv.Shutdown(ctx)
}

// setShutdownState - функция устанавливает флаг завершения приложения в указанное состояние
func (s *Server) setShutdownState(state bool) {
	s.shutdownState = state
}

// обогощает контекст идентификатором запроса
func (s *Server) middlewareRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.NewV4().String()
			r = r.WithContext(context.WithValue(r.Context(), RequestId, requestID))
			next.ServeHTTP(w, r)
		},
	)
}

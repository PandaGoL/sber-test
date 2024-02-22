package v1

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"sber-test/internal/services/deposit"
	"testing"

	"github.com/gorilla/mux"
)

func TestHandleSlothfulMessage(t *testing.T) {
	mainR := mux.NewRouter()

	r := &Router{
		mainRouter: mainR,
		router:     mainR.PathPrefix("/v1").Subrouter(),
		Service:    deposit.New(),
	}

	testTable := []struct {
		req            string
		expectedStatus int
	}{
		{
			req:            `{"date":"31.01.2021","periods":3,"amount":10000,"rate":6}`,
			expectedStatus: http.StatusOK,
		},
		{
			req:            "",
			expectedStatus: http.StatusBadRequest,
		},
		{
			req:            `{"date":"31/01/2021","periods":3,"amount":10000,"rate":6}`,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range testTable {

		wr := httptest.NewRecorder()

		req := httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBufferString(tt.req))

		r.CalculateDeposit(wr, req)

		if wr.Code != tt.expectedStatus {
			t.Errorf("got HTTP status code %d, expected: %d 200", wr.Code, tt.expectedStatus)
		}
	}

}

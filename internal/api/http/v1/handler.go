package v1

import (
	"encoding/json"
	"net/http"

	"sber-test/internal/services/deposit/types"

	uuid "github.com/satori/go.uuid"
)

func (rout *Router) CalculateDeposit(w http.ResponseWriter, r *http.Request) {
	requestId, ok := r.Context().Value("request_id").(string)
	if !ok {
		requestId = uuid.NewV4().String()
	}
	rec := &types.DepositeRequest{}
	err := json.NewDecoder(r.Body).Decode(&rec)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := rout.Service.CalculateDeposit(requestId, rec)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusBadRequest)
		resrError := &types.ErrorResponse{}
		resrError.Err = err.Error()
		data, err := json.Marshal(resrError)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)

		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(resp)
}

package deposit

import (
	"fmt"
	"sber-test/internal/services/deposit/types"
	"sber-test/pkg/utils"
	"time"

	"github.com/sirupsen/logrus"
)

//go:generate mockery --name=DepositServicer
type DepositServicer interface {
	CalculateDeposit(requestId string, req *types.DepositeRequest) ([]byte, error)
}

func (ds *DepositService) CalculateDeposit(requestId string, req *types.DepositeRequest) ([]byte, error) {

	logrus.WithField("request_id", requestId).WithFields(req.ToMap()).Info("Start CalculateDeposit request: ")

	err := ds.v.Struct(req)
	if err != nil {
		logrus.Errorf("request_id: %s Error in CalculateDeposit: %s", requestId, err)
		return nil, types.ErrValidation
	}

	resp := calculateDeposit(req)

	logrus.WithField("request_id", requestId).Info("Finish CalculateDeposit request")

	return resp, nil
}

func calculateDeposit(req *types.DepositeRequest) []byte {

	str := "{"

	var tmpAmont = float64(req.Amount)

	startPeriod, _ := time.Parse(types.Layout, req.Date)

	for i := 0; i < req.Periods; i++ {
		tmpAmont = calcAmount(tmpAmont, req.Rate)
		period := utils.AddMonth(startPeriod, i)
		if i < req.Periods-1 {
			str += fmt.Sprintf("\"%s\": %v,", period.Format(types.Layout), utils.RoundFloat(float64(tmpAmont), 2))
		}
		if i == req.Periods-1 {
			str += fmt.Sprintf("\"%s\": %v}", period.Format(types.Layout), utils.RoundFloat(float64(tmpAmont), 2))
		}
	}

	return []byte(str)
}

func calcAmount(amount float64, rate float64) float64 {
	return amount * (1 + rate/12/100)
}

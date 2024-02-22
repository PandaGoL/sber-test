package deposit

import (
	"github.com/go-playground/validator/v10"
)

type DepositService struct {
	v *validator.Validate
}

func New() *DepositService {
	us := &DepositService{
		v: validator.New(),
	}
	return us
}

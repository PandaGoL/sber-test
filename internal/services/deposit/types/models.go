package types

type DepositeRequest struct {
	Date    string  `json:"date" validate:"datetime=02.01.2006"`
	Periods int     `json:"periods" validate:"min=1,max=59"`
	Amount  int     `json:"amount" validate:"min=10000,max=2999999"`
	Rate    float64 `json:"rate" validate:"min=1,max=7"`
}

func (d *DepositeRequest) ToMap() map[string]interface{} {
	fields := make(map[string]interface{})
	fields["date"] = d.Date
	fields["periods"] = d.Periods
	fields["amount"] = d.Amount
	fields["rate"] = d.Rate

	return fields
}

type ErrorResponse struct {
	Err string `json:"error"`
}

package types

import "errors"

var (
	ErrValidation = errors.New("validation error")
	ErrParseDate  = errors.New("date parse error")
)

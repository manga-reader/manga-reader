package utils

import (
	"errors"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrUnsupportedEnv = errors.New("unsupported environment")
	ErrStrategy       = errors.New("sth wrong in trading strategy")
	ErrPriceInBound   = errors.New("price difference not big enough (in price bound)")
	ErrInternal       = errors.New("internal error")
	ErrLackOfBalance  = errors.New("lack of balance")
)

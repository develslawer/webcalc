package calculation

import "errors"

var (
	ErrUnknownOperator   = errors.New("unknown operator")
	ErrInvalidExpression = errors.New("invalid expression")
	ErrStackOverflow     = errors.New("stack overflow")
	ErrDivisionByZero    = errors.New("division by zero")
)

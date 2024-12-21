package application

import "errors"

var (
	ErrInvalidExpression = errors.New("Expression is not valid")
	ErrInternalServer    = errors.New("Internal server error")
)

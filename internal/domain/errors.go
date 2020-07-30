package domain

import "errors"

var (
	// ErrInternalServerError error internal server error
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound error not found
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict error conflict
	ErrConflict = errors.New("Your item alreadt exist")
	// ErrBadParamInput error bad param input
	ErrBadParamInput = errors.New("Given param is not valid")
)

package govern

import "errors"

var (
	ErrRequired       = errors.New("field is required")
	ErrMinLength      = errors.New("length is lower than min")
	ErrMaxLength      = errors.New("length is greater than max")
	ErrValueNotInEnum = errors.New("value is not in enum")
)

type (
	Error struct {
		error
		Field string
		Value any
	}
)

func newErr(inner error, field string, value any) *Error {
	return &Error{
		error: inner,
		Field: field,
		Value: value,
	}
}

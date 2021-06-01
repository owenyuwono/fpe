package fpe

// Error is the error type for fpe package errors
type Error string

func (e Error) Error() string {
	return string(e)
}

// ErrNegativeArgs error is returned when modulus argument is a negative number
var ErrNegativeArgs Error = "negative numbers cannot be used as modulus"

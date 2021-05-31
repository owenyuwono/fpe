package fpe

// Error is the error type for fpe package errors
type Error string

func (e Error) Error() string {
	return string(e)
}

// ErrModTooSmall ...
var ErrModTooSmall Error = "modulus range is too small"

// ErrNegativeArgs ...
var ErrNegativeArgs Error = "negative numbers cannot be used as modulus"

package cerrors

import "errors"

var (
	ErrUserNil = errors.New("error user is Nil")
	ErrBodyNil = errors.New("error body is Nil")

	ErrUnmarshalData = errors.New("error decode data")
)

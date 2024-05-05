package errors

import "errors"

var (
	ErrInvalidRequest     = errors.New("invalid request")
	ErrDuplicatedEntity   = errors.New("duplicated entity")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEntityNotFound     = errors.New("entity not found")
)

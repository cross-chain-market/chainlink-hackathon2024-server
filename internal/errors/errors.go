package errors

import "errors"

var (
	ErrInvalidRequest                            = errors.New("invalid request")
	ErrDuplicatedEntity                          = errors.New("duplicated entity")
	ErrInvalidCredentials                        = errors.New("invalid credentials")
	ErrEntityNotFound                            = errors.New("entity not found")
	ErrListedAmountGreaterThanTotalAmount        = errors.New("listed amount greater than total amount")
	ErrCannotUpdateFiatPrice                     = errors.New("cannot update fiat price")
	ErrCannotUnlistGreaterAmountThanListedAmount = errors.New("cannot unlist amount greater than listed amount")
	ErrCannotBuyMoreThanListedAmount             = errors.New("cannot buy more than listed amount")
)

package blockchain

import "errors"

var (
	ErrInvalidKey             = errors.New("invalid key")
	ErrInvalidAddress         = errors.New("invalid address")
	ErrInvalidContractAddress = errors.New("invalid contract address")
)

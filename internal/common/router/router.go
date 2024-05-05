package router

import (
	"context"
	"fmt"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/common/validator"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/errors"
	"net/http"

	"github.com/ggicci/httpin"
)

func ParseInput[T any](ctx context.Context) (*T, error) {
	input, ok := ctx.Value(httpin.Input).(*T)
	if !ok {
		return input, errors.ErrInvalidRequest
	}

	return input, validator.Validator.Struct(input)
}

func AddCacheHeader(w http.ResponseWriter, seconds int) {
	w.Header().Set("Cache-Control", fmt.Sprintf("max-age=%d", seconds))
}

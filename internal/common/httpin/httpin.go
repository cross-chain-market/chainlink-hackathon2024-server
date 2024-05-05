package httpin

import (
	"errors"
	"fmt"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/common/response"
	"github.com/ggicci/httpin"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
)

var (
	uuidDecoderFunc   = httpin.DecoderFunc[string](uuidDecoder)
	stringDecoderFunc = httpin.DecoderFunc[string](stringDecoder)
)

func InitHTTPIn() {
	httpin.UseGorillaMux("path", mux.Vars)
	httpin.RegisterValueTypeDecoder[uuid.UUID](uuidDecoderFunc)
	httpin.RegisterValueTypeDecoder[*string](stringDecoderFunc)
	httpin.ReplaceDefaultErrorHandler(defaultErrorHandler)
}

func uuidDecoder(s string) (any, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to parse [%s] uuid from request: %w", s, err)
	}

	return id, nil
}

func stringDecoder(s string) (any, error) {
	if s == "" {
		return (*string)(nil), nil
	}

	return &s, nil
}

func defaultErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	var invalidFieldError *httpin.InvalidFieldError
	if errors.As(err, &invalidFieldError) {
		response.BadRequest(w, err)
		return
	}

	if r.Method == http.MethodPost || r.Method == http.MethodPut {
		response.BadRequest(w, errors.New("method should be POST or PUT in order to use httpin"))
		return
	}

	slog.ErrorContext(r.Context(), "failed to parse request", slog.String("error", err.Error()))
	response.InternalServerError(w, err)
	return
}

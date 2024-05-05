package response

import (
	"encoding/json"
	errs "github.com/cross-chain-market/chainlink-hackathon2024-server/internal/errors"
	"net/http"
)

func BadRequest(w http.ResponseWriter, err error) {
	writeError(w, err, http.StatusBadRequest)
	return
}

func InternalServerError(w http.ResponseWriter, err error) {
	writeError(w, err, http.StatusInternalServerError)
	return
}

func Unauthorized(w http.ResponseWriter) {
	writeError(w, errs.ErrInvalidCredentials, http.StatusUnauthorized)
	return
}

func Ok(w http.ResponseWriter, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(body)
	return
}

func writeError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	return
}

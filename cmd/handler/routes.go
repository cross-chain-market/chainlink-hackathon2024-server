package handler

import (
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/config"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/marketplace"
	"github.com/ggicci/httpin"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	healthGo "github.com/hellofresh/health-go"
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"net/http"
)

type (
	getUserByEmailRequest struct {
		Email string `in:"path=email"`
	}

	registerUserRequest struct {
		Body registerUserInput `in:"body=json"`
	}

	loginUserRequest struct {
		Body loginUserInput `in:"body=json"`
	}

	updateUserRequest struct {
		RewardID uuid.UUID         `in:"path=rewardId"`
		Body     registerUserInput `in:"body=json"`
	}

	deleteUserRequest struct {
		UserID uuid.UUID `in:"path=userId"`
	}

	registerUserInput struct {
		Email    string `json:"email" validate:"notblank"`
		Username string `json:"username" validate:"notblank"`
		Password string `json:"password" validate:"notblank"`
	}

	loginUserInput struct {
		Email    string `json:"email" validate:"notblank"`
		Password string `json:"password" validate:"notblank"`
	}
)

func InitRoutes(cfg *config.Config, marketplaceService *marketplace.Service) *http.Server {
	r := mux.NewRouter()

	r.HandleFunc("/health", healthGo.HandlerFunc).Methods(http.MethodGet)

	h := &marketplaceHandler{service: marketplaceService}

	v1Router := r.NewRoute().PathPrefix("/v1").Subrouter()

	// Users
	v1Router.Handle("/users/register", alice.New(httpin.NewInput(registerUserRequest{})).ThenFunc(h.registerUser)).Methods(http.MethodPost)
	v1Router.Handle("/users/login", alice.New(httpin.NewInput(loginUserRequest{})).ThenFunc(h.loginUser)).Methods(http.MethodPost)
	v1Router.Handle("/users/{email}", alice.New(httpin.NewInput(getUserByEmailRequest{})).ThenFunc(h.getUserByEmail)).Methods(http.MethodGet)
	v1Router.Handle("/users/{userId}", alice.New(httpin.NewInput(updateUserRequest{})).ThenFunc(h.updateUser)).Methods(http.MethodPut)
	v1Router.Handle("/users/{userId}", alice.New(httpin.NewInput(deleteUserRequest{})).ThenFunc(h.deleteUser)).Methods(http.MethodDelete)

	return &http.Server{
		Addr:    ":" + cfg.Profile.Port,
		Handler: cors.AllowAll().Handler(r),
	}
}

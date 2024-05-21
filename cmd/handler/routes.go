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

	createCollectionRequest struct {
		Body createCollectionInput `in:"body=json"`
	}

	getCollectionRequest struct {
		UserAddress  string `in:"path=userAddress"`
		CollectionID int64  `in:"path=collectionId"`
	}

	listItemsRequest struct {
		CollectionID int64        `in:"path=collectionId"`
		ItemID       int64        `in:"path=itemId"`
		Body         listingInput `in:"body=json"`
	}

	getListingsRequest struct {
		CollectionID *int64 `in:"query=collectionId"`

		// TODO: Add more optional filters like userID for example
	}

	buyItemsRequest struct {
		CollectionID int64         `in:"path=collectionId"`
		ItemID       int64         `in:"path=itemId"`
		Body         buyItemsInput `in:"body=json"`
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

	createCollectionInput struct {
		Name            string       `json:"name" validate:"notblank"`
		Description     string       `json:"description"`
		BaseHash        string       `json:"base_hash"`
		OwnerAddressHex string       `json:"owner_address_hex"`
		NetworkID       string       `json:"network_id"`
		ChainID         int64        `json:"chain_id"`
		Items           []itemsInput `json:"items"`
	}

	itemsInput struct {
		Name         string         `json:"name" validate:"notblank"`
		Description  string         `json:"description"`
		ImageID      string         `json:"image_id"`
		FiatPrice    float64        `json:"fiat_price"`
		TotalAmount  int64          `json:"total_amount"`
		ListedAmount int64          `json:"listed_amount"`
		Attributes   map[string]any `json:"attributes"`
	}

	listingInput struct {
		FiatPrice    float64 `json:"fiat_price"`
		ListedAmount int64   `json:"listed_amount"`
	}

	buyItemsInput struct {
		FromAddress string `json:"from_address"`
		ToAddress   string `json:"to_address"`
		Amount      int64  `json:"amount"`
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

	// Collections
	v1Router.Handle("/collections", alice.New(httpin.NewInput(createCollectionRequest{})).ThenFunc(h.createCollection)).Methods(http.MethodPost)
	v1Router.Handle("/users/{userAddress}/collections", alice.New(httpin.NewInput(getCollectionRequest{})).ThenFunc(h.getUserCollections)).Methods(http.MethodGet)
	v1Router.Handle("/users/{userAddress}/collections/{collectionId}", alice.New(httpin.NewInput(getCollectionRequest{})).ThenFunc(h.getCollection)).Methods(http.MethodGet)

	// List/Unlist items
	v1Router.Handle("/collections/{collectionId}/items/{itemId}/list", alice.New(httpin.NewInput(listItemsRequest{})).ThenFunc(h.listItems)).Methods(http.MethodPost)
	v1Router.Handle("/collections/{collectionId}/items/{itemId}/unlist", alice.New(httpin.NewInput(listItemsRequest{})).ThenFunc(h.unlistItems)).Methods(http.MethodPost)

	// Get listed items
	v1Router.Handle("/listings", alice.New(httpin.NewInput(getListingsRequest{})).ThenFunc(h.getListings)).Methods(http.MethodPost)

	// Buy listed items
	v1Router.Handle("/collections/{collectionId}/items/{itemId}/buy", alice.New(httpin.NewInput(buyItemsRequest{})).ThenFunc(h.buyItems)).Methods(http.MethodPost)

	return &http.Server{
		Addr:    ":" + cfg.Profile.Port,
		Handler: cors.AllowAll().Handler(r),
	}
}

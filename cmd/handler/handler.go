package handler

import (
	"errors"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/common/response"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/common/router"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/marketplace"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/marketplace/model"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type (
	marketplaceHandler struct {
		service *marketplace.Service
	}
)

func (h *marketplaceHandler) createCollection(w http.ResponseWriter, r *http.Request) {
	request, err := router.ParseInput[createCollectionRequest](r.Context())
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	if len(request.Body.Items) == 0 {
		response.BadRequest(w, errors.New("no items provided"))
	}

	collectionID := uuid.New()

	items := make([]*model.Item, 0, len(request.Body.Items))

	for _, item := range request.Body.Items {
		items = append(items, &model.Item{
			ID:            uuid.New(),
			CollectionID:  collectionID,
			Name:          item.Name,
			Description:   item.Description,
			BaseImagePath: item.BaseImagePath,
			FiatPrice:     item.FiatPrice,
			Address:       item.Address,
			TotalAmount:   item.TotalAmount,
			ListedAmount:  item.ListedAmount,
		})
	}

	collection := &model.Collection{
		ID:            collectionID,
		UserID:        request.UserID,
		Name:          request.Body.Name,
		Description:   request.Body.Description,
		BaseImagePath: request.Body.BaseImagePath,
		Items:         items,
	}

	result, err := h.service.CreateCollection(r.Context(), collection)
	if err != nil {
		response.InternalServerError(w, err)
		return
	}

	response.Ok(w, result)
	return
}

func (h *marketplaceHandler) registerUser(w http.ResponseWriter, r *http.Request) {
	request, err := router.ParseInput[registerUserRequest](r.Context())
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	user := &model.User{
		ID:        uuid.New(),
		Email:     request.Body.Email,
		Username:  request.Body.Username,
		Password:  request.Body.Password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	result, err := h.service.RegisterUser(r.Context(), user)
	if err != nil {
		response.InternalServerError(w, err)
		return
	}

	response.Ok(w, result)
	return
}

func (h *marketplaceHandler) loginUser(w http.ResponseWriter, r *http.Request) {
	request, err := router.ParseInput[loginUserRequest](r.Context())
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	success, err := h.service.LoginUser(r.Context(), request.Body.Email, request.Body.Password)
	if err != nil {
		response.InternalServerError(w, err)
		return
	}

	if success {
		response.Ok(w, "")
		return
	} else {
		response.Unauthorized(w)
		return
	}
}

func (h *marketplaceHandler) getUserByEmail(w http.ResponseWriter, r *http.Request) {
	request, err := router.ParseInput[getUserByEmailRequest](r.Context())
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	result, err := h.service.GetUserByEmail(r.Context(), request.Email)
	if err != nil {
		response.InternalServerError(w, err)
		return
	}

	response.Ok(w, result)
	return
}

func (h *marketplaceHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	// TODO

}

func (h *marketplaceHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	// TODO

}

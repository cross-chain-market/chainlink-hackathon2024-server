package handler

import (
	"errors"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/common/response"
	"github.com/cross-chain-market/chainlink-hackathon2024-server/internal/common/router"
	errs "github.com/cross-chain-market/chainlink-hackathon2024-server/internal/errors"
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
		return
	}

	items := make([]*model.Item, 0, len(request.Body.Items))

	for _, item := range request.Body.Items {

		if item.ListedAmount > item.TotalAmount {
			response.BadRequest(w, errors.New("listed amount cannot be greater than total amount"))
			return
		}

		items = append(items, &model.Item{
			Name:         item.Name,
			Description:  item.Description,
			ImageID:      item.ImageID,
			FiatPrice:    item.FiatPrice,
			TotalAmount:  item.TotalAmount,
			ListedAmount: item.ListedAmount,
			Attributes:   item.Attributes,
			CreatedAt:    time.Now().UTC(),
		})
	}

	collection := &model.Collection{
		OwnerAddressHex: request.UserAddress,
		Name:            request.Body.Name,
		Description:     request.Body.Description,
		BaseHash:        request.Body.BaseHash,
		NetworkID:       request.Body.NetworkID,
		ChainID:         request.Body.ChainID,
		Status:          model.NotDeployedStatus,
		Items:           items,
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
	}

	result, err := h.service.CreateCollection(r.Context(), collection)
	if err != nil {
		response.InternalServerError(w, err)
		return
	}

	response.Ok(w, result)
	return
}

func (h *marketplaceHandler) getCollection(w http.ResponseWriter, r *http.Request) {
	request, err := router.ParseInput[getCollectionRequest](r.Context())
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	result, err := h.service.GetCollection(r.Context(), request.CollectionID)
	if err != nil {
		response.InternalServerError(w, err)
		return
	}

	response.Ok(w, result)
	return
}

func (h *marketplaceHandler) getUserCollections(w http.ResponseWriter, r *http.Request) {
	request, err := router.ParseInput[getCollectionRequest](r.Context())
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	result, err := h.service.GetUserCollections(r.Context(), request.UserAddress)
	if err != nil {
		response.InternalServerError(w, err)
		return
	}

	response.Ok(w, result)
	return
}

func (h *marketplaceHandler) listItems(w http.ResponseWriter, r *http.Request) {
	request, err := router.ParseInput[listItemsRequest](r.Context())
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	if request.Body.ListedAmount <= 0 || request.Body.FiatPrice <= 0 {
		response.BadRequest(w, errs.ErrInvalidRequest)
		return
	}

	result, err := h.service.ListItem(r.Context(), request.CollectionID, request.ItemID, request.Body.ListedAmount, request.Body.FiatPrice)
	if err != nil {
		response.InternalServerError(w, err)
		return
	}

	response.Ok(w, result)
	return
}

func (h *marketplaceHandler) unlistItems(w http.ResponseWriter, r *http.Request) {
	request, err := router.ParseInput[listItemsRequest](r.Context())
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	if request.Body.ListedAmount <= 0 {
		response.BadRequest(w, errs.ErrInvalidRequest)
		return
	}

	result, err := h.service.UnlistItem(r.Context(), request.CollectionID, request.ItemID, request.Body.ListedAmount)
	if err != nil {
		response.InternalServerError(w, err)
		return
	}

	response.Ok(w, result)
	return
}

func (h *marketplaceHandler) getListings(w http.ResponseWriter, r *http.Request) {
	request, err := router.ParseInput[getListingsRequest](r.Context())
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	result, err := h.service.GetListings(r.Context(), request.CollectionID)
	if err != nil {
		response.InternalServerError(w, err)
		return
	}

	response.Ok(w, result)
	return
}

func (h *marketplaceHandler) buyItems(w http.ResponseWriter, r *http.Request) {
	request, err := router.ParseInput[buyItemsRequest](r.Context())
	if err != nil {
		response.BadRequest(w, err)
		return
	}

	if request.Body.Amount <= 0 {
		response.BadRequest(w, errs.ErrInvalidRequest)
		return
	}

	result, err := h.service.BuyItems(r.Context(), request.CollectionID, request.ItemID, request.Body.Amount, request.Body.FromAddress, request.Body.ToAddress)
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

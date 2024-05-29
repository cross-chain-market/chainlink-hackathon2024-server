package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type (
	User struct {
		bun.BaseModel `bun:"table:users"`

		ID        uuid.UUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
		Email     string     `bun:"email,notnull" json:"email"`
		Username  string     `bun:"username,notnull" json:"username"`
		Password  string     `bun:"password" json:"-"`
		CreatedAt time.Time  `bun:"created_at,default:current_timestamp" json:"created_at"`
		UpdatedAt time.Time  `bun:"updated_at,default:current_timestamp" json:"updated_at"`
		DeletedAt *time.Time `bun:"deleted_at,soft_delete" json:"-"`
	}

	Collection struct {
		bun.BaseModel `bun:"table:collections"`

		ID              int64      `bun:"id,pk,autoincrement" json:"id"`
		OwnerAddressHex string     `bun:"owner_address_hex,notnull" json:"owner_address_hex"`
		Name            string     `bun:"name" json:"name"`
		Description     string     `bun:"description" json:"description"`
		BaseHash        string     `bun:"base_hash" json:"base_hash"`
		Address         string     `bun:"address" json:"address"`
		NetworkID       string     `bun:"network_id,notnull" json:"network_id"`
		ChainID         int64      `bun:"chain_id,notnull" json:"chain_id"`
		Items           []*Item    `bun:"rel:has-many,join:id=collection_id" json:"items"`
		Status          string     `bun:"status,notnull" json:"status"`
		CreatedAt       time.Time  `bun:"created_at,default:current_timestamp" json:"created_at"`
		UpdatedAt       time.Time  `bun:"updated_at,default:current_timestamp" json:"updated_at"`
		DeletedAt       *time.Time `bun:"deleted_at,soft_delete" json:"-"`
	}

	Item struct {
		bun.BaseModel `bun:"table:items"`

		ID                int64          `bun:"id,pk,autoincrement" json:"id"`
		CollectionID      int64          `bun:"collection_id,notnull" json:"collection_id"`
		Name              string         `bun:"name" json:"name"`
		Description       string         `bun:"description" json:"description"`
		ImageID           int64          `bun:"image_id" json:"image_id"`
		FiatPrice         float64        `bun:"fiat_price" json:"fiat_price"`
		TotalAmount       int64          `bun:"total_amount" json:"total_amount"`
		ListedAmount      int64          `bun:"listed_amount" json:"listed_amount"`
		Attributes        map[string]any `bun:"attributes" json:"attributes"`
		CreatedAt         time.Time      `bun:"created_at,default:current_timestamp" json:"created_at"`
		DeletedAt         *time.Time     `bun:"deleted_at,soft_delete" json:"-"`
		NetworkID         *string        `bun:"-" json:"network_id,omitempty"`
		ChainID           *int64         `bun:"-" json:"chain_id,omitempty"`
		BaseHash          *string        `bun:"-" json:"base_hash,omitempty"`
		CollectionAddress *string        `bun:"-" json:"address,omitempty"`
	}
)

const (
	NotDeployedStatus = "NOT_DEPLOYED"
	PendingTXStatus   = "PENDING_TX"
	DeployedStatus    = "DEPLOYED"
)

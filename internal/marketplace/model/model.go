package model

import (
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
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

		ID            uuid.UUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
		UserID        uuid.UUID  `bun:"user_id,notnull" json:"user_id"`
		Name          string     `bun:"name" json:"name"`
		Description   string     `bun:"description" json:"description"`
		BaseImagePath string     `bun:"base_image_path" json:"base_image_path"`
		Items         []*Item    `bun:"rel:has-many,join:id=collection_id" json:"items"`
		CreatedAt     time.Time  `bun:"created_at,default:current_timestamp" json:"created_at"`
		UpdatedAt     time.Time  `bun:"updated_at,default:current_timestamp" json:"updated_at"`
		DeletedAt     *time.Time `bun:"deleted_at,soft_delete" json:"-"`
	}

	Item struct {
		bun.BaseModel `bun:"table:items"`

		ID            uuid.UUID  `bun:"id,pk,type:uuid,default:uuid_generate_v4()" json:"id"`
		CollectionID  uuid.UUID  `bun:"collection_id,notnull" json:"collection_id"`
		Name          string     `bun:"name" json:"name"`
		Description   string     `bun:"description" json:"description"`
		BaseImagePath string     `bun:"base_image_path" json:"base_image_path"`
		FiatPrice     float64    `bun:"fiat_price" json:"fiat_price"`
		Address       string     `bun:"address" json:"address"`
		TotalAmount   int64      `bun:"total_amount" json:"total_amount"`
		ListedAmount  int64      `bun:"listed_amount" json:"listed_amount"`
		CreatedAt     time.Time  `bun:"created_at,default:current_timestamp" json:"created_at"`
		DeletedAt     *time.Time `bun:"deleted_at,soft_delete" json:"-"`
	}
)

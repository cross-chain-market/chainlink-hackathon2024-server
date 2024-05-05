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
		CreatedAt time.Time  `bun:"created_at,default:current_timestamp" json:"createdAt"`
		UpdatedAt time.Time  `bun:"updated_at,default:current_timestamp" json:"updatedAt"`
		DeletedAt *time.Time `bun:"deleted_at,soft_delete" json:"-"`
	}
)

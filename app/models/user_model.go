package models

import (
	"time"

	"github.com/STEEDUj2kb/pkg/utils"
	"github.com/STEEDUj2kb/platform/ent/user"
	"github.com/google/uuid"
)

// User struct to describe User object.
type User struct {
	Id           int           `db:"id" json:"-"`
	UUID         uuid.UUID     `db:"uuid" json:"uuid" validate:"required,uuid"`
	CreatedAt    time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at" json:"updated_at"`
	Name         string        `db:"name" json:"name" validate:"required,lte=255"`
	Email        string        `db:"email" json:"email" validate:"required,email,lte=255"`
	PasswordHash string        `db:"password_hash" json:"-" validate:"required,lte=255"`
	UserStatus   int           `db:"user_status" json:"user_status"`
	UserRole     user.UserRole `db:"user_role" json:"user_role"`
}

func (user *User) ApplyData(fromModel interface{}) {
	utils.ApplyData(user, fromModel)
}

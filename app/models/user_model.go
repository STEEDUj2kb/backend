package models

// User struct to describe User object.
type User struct {
	Name  string `db:"name" json:"name" validate:"required,lte=255"`
	Email string `db:"email" json:"email" validate:"required,email,lte=255"`
}

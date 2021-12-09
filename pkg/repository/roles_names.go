package repository

// UserRole defines the type for the "user_role" enum field.
type UserRole string

// UserRole values.
const (
	// AdminRoleName const for admin role.
	Admin UserRole = "admin"
	// UserRoleName const for user role.
	User UserRole = "user"
)

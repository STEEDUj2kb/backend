package utils

import (
	"fmt"

	"github.com/STEEDUj2kb/pkg/repository"
)

// VerifyRole func for verifying a given role.
func VerifyRole(role repository.UserRole) (repository.UserRole, error) {
	// Switch given role.
	switch role {
	case repository.Admin:
		// Nothing to do, verified successfully.
	case repository.User:
		// Nothing to do, verified successfully.
	default:
		// Return error message.
		return "", fmt.Errorf("role '%v' does not exist", role)
	}

	return role, nil
}

package utils

import (
	"fmt"

	"github.com/STEEDUj2kb/pkg/repository"
)

// GetCredentialsByRole func for getting credentials from a role name.
func GetCredentialsByRole(role repository.UserRole) ([]string, error) {
	// Define credentials variable.
	var credentials []string

	// Switch given role.
	switch role {
	case repository.Admin:
		// Admin credentials (all access).
		credentials = []string{
			repository.StudyCreateCredential,
			repository.StudyUpdateCredential,
			repository.StudyDeleteCredential,
		}
	case repository.User:
		// Simple user credentials (only book creation).
		credentials = []string{
			repository.StudyCreateCredential,
		}
	default:
		// Return error message.
		return nil, fmt.Errorf("role '%v' does not exist", role)
	}

	return credentials, nil
}

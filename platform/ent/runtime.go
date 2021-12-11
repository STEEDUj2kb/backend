// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/STEEDUj2kb/platform/ent/schema"
	"github.com/STEEDUj2kb/platform/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUUID is the schema descriptor for uuid field.
	userDescUUID := userFields[0].Descriptor()
	// user.DefaultUUID holds the default value on creation for the uuid field.
	user.DefaultUUID = userDescUUID.Default.(func() uuid.UUID)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[1].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[2].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescUserStatus is the schema descriptor for user_status field.
	userDescUserStatus := userFields[5].Descriptor()
	// user.DefaultUserStatus holds the default value on creation for the user_status field.
	user.DefaultUserStatus = userDescUserStatus.Default.(int)
	// user.UserStatusValidator is a validator for the "user_status" field. It is called by the builders before save.
	user.UserStatusValidator = userDescUserStatus.Validators[0].(func(int) error)
}

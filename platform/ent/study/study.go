// Code generated by entc, DO NOT EDIT.

package study

import (
	"time"
)

const (
	// Label holds the string label denoting the study type in the database.
	Label = "study"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStartedAt holds the string denoting the started_at field in the database.
	FieldStartedAt = "started_at"
	// FieldEndedAt holds the string denoting the ended_at field in the database.
	FieldEndedAt = "ended_at"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// EdgePlanner holds the string denoting the planner edge name in mutations.
	EdgePlanner = "planner"
	// EdgeDgoals holds the string denoting the dgoals edge name in mutations.
	EdgeDgoals = "dgoals"
	// EdgeWgoals holds the string denoting the wgoals edge name in mutations.
	EdgeWgoals = "wgoals"
	// Table holds the table name of the study in the database.
	Table = "studies"
	// PlannerTable is the table that holds the planner relation/edge.
	PlannerTable = "studies"
	// PlannerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	PlannerInverseTable = "users"
	// PlannerColumn is the table column denoting the planner relation/edge.
	PlannerColumn = "user_studies"
	// DgoalsTable is the table that holds the dgoals relation/edge.
	DgoalsTable = "daily_gaols"
	// DgoalsInverseTable is the table name for the DailyGaol entity.
	// It exists in this package in order to avoid circular dependency with the "dailygaol" package.
	DgoalsInverseTable = "daily_gaols"
	// DgoalsColumn is the table column denoting the dgoals relation/edge.
	DgoalsColumn = "study_dgoals"
	// WgoalsTable is the table that holds the wgoals relation/edge.
	WgoalsTable = "weekly_gaols"
	// WgoalsInverseTable is the table name for the WeeklyGaol entity.
	// It exists in this package in order to avoid circular dependency with the "weeklygaol" package.
	WgoalsInverseTable = "weekly_gaols"
	// WgoalsColumn is the table column denoting the wgoals relation/edge.
	WgoalsColumn = "study_wgoals"
)

// Columns holds all SQL columns for study fields.
var Columns = []string{
	FieldID,
	FieldStartedAt,
	FieldEndedAt,
	FieldContent,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "studies"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_studies",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultStartedAt holds the default value on creation for the "started_at" field.
	DefaultStartedAt func() time.Time
	// DefaultEndedAt holds the default value on creation for the "ended_at" field.
	DefaultEndedAt func() time.Time
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
)

// Code generated by entc, DO NOT EDIT.

package weeklygaol

const (
	// Label holds the string label denoting the weeklygaol type in the database.
	Label = "weekly_gaol"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldGoal holds the string denoting the goal field in the database.
	FieldGoal = "goal"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// FieldNth holds the string denoting the nth field in the database.
	FieldNth = "nth"
	// EdgeDgaols holds the string denoting the dgaols edge name in mutations.
	EdgeDgaols = "dgaols"
	// EdgeStudy holds the string denoting the study edge name in mutations.
	EdgeStudy = "study"
	// Table holds the table name of the weeklygaol in the database.
	Table = "weekly_gaols"
	// DgaolsTable is the table that holds the dgaols relation/edge.
	DgaolsTable = "daily_gaols"
	// DgaolsInverseTable is the table name for the DailyGaol entity.
	// It exists in this package in order to avoid circular dependency with the "dailygaol" package.
	DgaolsInverseTable = "daily_gaols"
	// DgaolsColumn is the table column denoting the dgaols relation/edge.
	DgaolsColumn = "weekly_gaol_dgaols"
	// StudyTable is the table that holds the study relation/edge.
	StudyTable = "weekly_gaols"
	// StudyInverseTable is the table name for the Study entity.
	// It exists in this package in order to avoid circular dependency with the "study" package.
	StudyInverseTable = "studies"
	// StudyColumn is the table column denoting the study relation/edge.
	StudyColumn = "study_wgoals"
)

// Columns holds all SQL columns for weeklygaol fields.
var Columns = []string{
	FieldID,
	FieldGoal,
	FieldScore,
	FieldNth,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "weekly_gaols"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"study_wgoals",
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
	// DefaultScore holds the default value on creation for the "score" field.
	DefaultScore int
)
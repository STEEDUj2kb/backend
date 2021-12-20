// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/STEEDUj2kb/platform/ent/daily_gaol"
	"github.com/STEEDUj2kb/platform/ent/study"
	"github.com/STEEDUj2kb/platform/ent/weekly_gaol"
)

// Daily_Gaol is the model entity for the Daily_Gaol schema.
type Daily_Gaol struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Todo holds the value of the "todo" field.
	Todo string `json:"todo,omitempty"`
	// Done holds the value of the "done" field.
	Done bool `json:"done,omitempty"`
	// IsRemoved holds the value of the "is_removed" field.
	IsRemoved bool `json:"is_removed,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the Daily_GaolQuery when eager-loading is set.
	Edges              Daily_GaolEdges `json:"edges"`
	study_dgoals       *int
	weekly_gaol_dgaols *int
}

// Daily_GaolEdges holds the relations/edges for other nodes in the graph.
type Daily_GaolEdges struct {
	// Study holds the value of the study edge.
	Study *Study `json:"study,omitempty"`
	// Wgoal holds the value of the wgoal edge.
	Wgoal *Weekly_Gaol `json:"wgoal,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// StudyOrErr returns the Study value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Daily_GaolEdges) StudyOrErr() (*Study, error) {
	if e.loadedTypes[0] {
		if e.Study == nil {
			// The edge study was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: study.Label}
		}
		return e.Study, nil
	}
	return nil, &NotLoadedError{edge: "study"}
}

// WgoalOrErr returns the Wgoal value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e Daily_GaolEdges) WgoalOrErr() (*Weekly_Gaol, error) {
	if e.loadedTypes[1] {
		if e.Wgoal == nil {
			// The edge wgoal was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: weekly_gaol.Label}
		}
		return e.Wgoal, nil
	}
	return nil, &NotLoadedError{edge: "wgoal"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Daily_Gaol) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case daily_gaol.FieldDone, daily_gaol.FieldIsRemoved:
			values[i] = new(sql.NullBool)
		case daily_gaol.FieldID:
			values[i] = new(sql.NullInt64)
		case daily_gaol.FieldTodo:
			values[i] = new(sql.NullString)
		case daily_gaol.FieldCreatedAt, daily_gaol.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case daily_gaol.ForeignKeys[0]: // study_dgoals
			values[i] = new(sql.NullInt64)
		case daily_gaol.ForeignKeys[1]: // weekly_gaol_dgaols
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Daily_Gaol", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Daily_Gaol fields.
func (dg *Daily_Gaol) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case daily_gaol.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dg.ID = int(value.Int64)
		case daily_gaol.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				dg.CreatedAt = value.Time
			}
		case daily_gaol.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				dg.UpdatedAt = value.Time
			}
		case daily_gaol.FieldTodo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field todo", values[i])
			} else if value.Valid {
				dg.Todo = value.String
			}
		case daily_gaol.FieldDone:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field done", values[i])
			} else if value.Valid {
				dg.Done = value.Bool
			}
		case daily_gaol.FieldIsRemoved:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_removed", values[i])
			} else if value.Valid {
				dg.IsRemoved = value.Bool
			}
		case daily_gaol.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field study_dgoals", value)
			} else if value.Valid {
				dg.study_dgoals = new(int)
				*dg.study_dgoals = int(value.Int64)
			}
		case daily_gaol.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field weekly_gaol_dgaols", value)
			} else if value.Valid {
				dg.weekly_gaol_dgaols = new(int)
				*dg.weekly_gaol_dgaols = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryStudy queries the "study" edge of the Daily_Gaol entity.
func (dg *Daily_Gaol) QueryStudy() *StudyQuery {
	return (&Daily_GaolClient{config: dg.config}).QueryStudy(dg)
}

// QueryWgoal queries the "wgoal" edge of the Daily_Gaol entity.
func (dg *Daily_Gaol) QueryWgoal() *WeeklyGaolQuery {
	return (&Daily_GaolClient{config: dg.config}).QueryWgoal(dg)
}

// Update returns a builder for updating this Daily_Gaol.
// Note that you need to call Daily_Gaol.Unwrap() before calling this method if this Daily_Gaol
// was returned from a transaction, and the transaction was committed or rolled back.
func (dg *Daily_Gaol) Update() *DailyGaolUpdateOne {
	return (&Daily_GaolClient{config: dg.config}).UpdateOne(dg)
}

// Unwrap unwraps the Daily_Gaol entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dg *Daily_Gaol) Unwrap() *Daily_Gaol {
	tx, ok := dg.config.driver.(*txDriver)
	if !ok {
		panic("ent: Daily_Gaol is not a transactional entity")
	}
	dg.config.driver = tx.drv
	return dg
}

// String implements the fmt.Stringer.
func (dg *Daily_Gaol) String() string {
	var builder strings.Builder
	builder.WriteString("Daily_Gaol(")
	builder.WriteString(fmt.Sprintf("id=%v", dg.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(dg.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(dg.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", todo=")
	builder.WriteString(dg.Todo)
	builder.WriteString(", done=")
	builder.WriteString(fmt.Sprintf("%v", dg.Done))
	builder.WriteString(", is_removed=")
	builder.WriteString(fmt.Sprintf("%v", dg.IsRemoved))
	builder.WriteByte(')')
	return builder.String()
}

// Daily_Gaols is a parsable slice of Daily_Gaol.
type Daily_Gaols []*Daily_Gaol

func (dg Daily_Gaols) config(cfg config) {
	for _i := range dg {
		dg[_i].config = cfg
	}
}

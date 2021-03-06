// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DailyGaolsColumns holds the columns for the "daily_gaols" table.
	DailyGaolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "todo", Type: field.TypeString},
		{Name: "done", Type: field.TypeBool, Default: false},
		{Name: "is_removed", Type: field.TypeBool, Default: false},
		{Name: "study_dgoals", Type: field.TypeInt, Nullable: true},
		{Name: "weekly_gaol_dgaols", Type: field.TypeInt, Nullable: true},
	}
	// DailyGaolsTable holds the schema information for the "daily_gaols" table.
	DailyGaolsTable = &schema.Table{
		Name:       "daily_gaols",
		Columns:    DailyGaolsColumns,
		PrimaryKey: []*schema.Column{DailyGaolsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "daily_gaols_studies_dgoals",
				Columns:    []*schema.Column{DailyGaolsColumns[6]},
				RefColumns: []*schema.Column{StudiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "daily_gaols_weekly_gaols_dgaols",
				Columns:    []*schema.Column{DailyGaolsColumns[7]},
				RefColumns: []*schema.Column{WeeklyGaolsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StudiesColumns holds the columns for the "studies" table.
	StudiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "started_at", Type: field.TypeTime},
		{Name: "ended_at", Type: field.TypeTime, Nullable: true},
		{Name: "content", Type: field.TypeString},
		{Name: "user_studies", Type: field.TypeInt, Nullable: true},
	}
	// StudiesTable holds the schema information for the "studies" table.
	StudiesTable = &schema.Table{
		Name:       "studies",
		Columns:    StudiesColumns,
		PrimaryKey: []*schema.Column{StudiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "studies_users_studies",
				Columns:    []*schema.Column{StudiesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uuid", Type: field.TypeUUID},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "user_status", Type: field.TypeInt, Default: 1},
		{Name: "user_role", Type: field.TypeEnum, Enums: []string{"admin", "user"}, Default: "user"},
		{Name: "password_hash", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// WeeklyGaolsColumns holds the columns for the "weekly_gaols" table.
	WeeklyGaolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "goal", Type: field.TypeString},
		{Name: "score", Type: field.TypeInt, Nullable: true, Default: 0},
		{Name: "nth", Type: field.TypeInt},
		{Name: "study_wgoals", Type: field.TypeInt, Nullable: true},
	}
	// WeeklyGaolsTable holds the schema information for the "weekly_gaols" table.
	WeeklyGaolsTable = &schema.Table{
		Name:       "weekly_gaols",
		Columns:    WeeklyGaolsColumns,
		PrimaryKey: []*schema.Column{WeeklyGaolsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "weekly_gaols_studies_wgoals",
				Columns:    []*schema.Column{WeeklyGaolsColumns[4]},
				RefColumns: []*schema.Column{StudiesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DailyGaolsTable,
		StudiesTable,
		UsersTable,
		WeeklyGaolsTable,
	}
)

func init() {
	DailyGaolsTable.ForeignKeys[0].RefTable = StudiesTable
	DailyGaolsTable.ForeignKeys[1].RefTable = WeeklyGaolsTable
	StudiesTable.ForeignKeys[0].RefTable = UsersTable
	WeeklyGaolsTable.ForeignKeys[0].RefTable = StudiesTable
}

package models

import (
	"time"

	"github.com/STEEDUj2kb/pkg/utils"
)

type Study struct {
	Id        int       `db:"id" json:"-"`
	Content   string    `db:"content" json:"content"`
	StartedAt time.Time `db:"started_at" json:"started_at"`
	EndedAt   time.Time `db:"ended_at" json:"ended_at"`
}

func (user *Study) ApplyData(fromModel interface{}) {
	utils.ApplyData(user, fromModel)
}

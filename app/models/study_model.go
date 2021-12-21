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

func (study *Study) ApplyData(fromModel interface{}) {
	utils.ApplyData(study, fromModel)
}

type WeeklyGaol struct {
	ID    int    `db:"id" json:"-"`
	Goal  string `db:"goal" json:"goal" validate:"required,lte=255"`
	Score int    `db:"score" json:"score"`
	Nth   int    `db:"nth" json:"nth"`
}

func (wgoal *WeeklyGaol) ApplyData(fromModel interface{}) {
	utils.ApplyData(wgoal, fromModel)
}

package data

import (
	"github.com/google/uuid"
)

type GYMIMAGE struct {
	Photo_ID uuid.UUID `json:"photo_id" gorm:"column:photo_id;type:uniqueidentifier;default:NEWID();primaryKey"`
	Gym_ID   uuid.UUID `json:"gym_id" gorm:"column:gym_id;type:uniqueidentifier"`
	Link     string    `json:"link" gorm:"type:varchar(300);column:link"`
}

func (GYMIMAGE) TableName() string {
	return "GYMIMAGES"
}

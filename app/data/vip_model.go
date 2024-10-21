package data

import (
	"time"

	"github.com/google/uuid"
)

type VIP struct {
	ID             uuid.UUID `json:"id" gorm:"type:uniqueidentifier;default:NEWID();primaryKey;column:id"`
	Is_VIP         bool      `json:"is_vip" gorm:"column:is_vip;type:bit"`
	VIP_start_date time.Time `json:"vip_start_date" gorm:"type:datetime;column:vip_start_date"`
	VIP_last_day   time.Time `json:"vip_last_day" gorm:"type:datetime;column:vip_last_day"`
	User_ID        uuid.UUID `json:"user_id" gorm:"type:uniqueidentifier;column:user_id"`
}

func (VIP) TableName() string {
	return "VIPS"
}

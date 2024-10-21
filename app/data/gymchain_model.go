package data

import (
	//"fmt"
	//"time"

	"github.com/google/uuid"
	//"golang.org/x/crypto/bcrypt"
)

type GYMCHAIN struct {
	ID             uuid.UUID `json:"id"  gorm:"column:id;type:uniqueidentifier;default:NEWID();primaryKey"`
	Name           string    `json:"name" gorm:"column:name;type:varchar(255);not null;"`
	Logo_link      string    `json:"logo_link" gorm:"type:varchar(255);column:logo_link"`
	Website        string    `json:"website" gorm:"type:varchar(255);column:website"`
	Description    string    `json:"description" gorm:"type:text;column:description"`
	Instagram_user string    `json:"instagram_user" gorm:"type:varchar(100);column:instagram_user"`
	Instagram_link string    `json:"instagram_link" gorm:"type:varchar(255);column:instagram_link"`
	Facebook_user  string    `json:"facebook_user" gorm:"type:varchar(100);column:facebook_user"`
	Facebook_link  string    `json:"facebook_link" gorm:"type:varchar(255);column:facebook_link"`

	Gyms []GYM `gorm:"foreignKey:gymchain_id"`
}

func (GYMCHAIN) TableName() string {
	return "GYMCHAINS"
}

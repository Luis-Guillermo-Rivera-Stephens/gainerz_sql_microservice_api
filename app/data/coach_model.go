package data

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type COACH struct {
	ID       uuid.UUID `json:"id"  gorm:"column:id;type:uniqueidentifier;default:NEWID();primaryKey"`
	Name     string    `json:"name" gorm:"column:name;type:varchar(255);not null;"`
	Email    string    `json:"email" gorm:"column:email;type:varchar(255);uniqueIndex"`
	Password string    `json:"password" gorm:"column:password;type:varchar(255);not null"`

	Description    string `json:"description" gorm:"column:description;type:varchar(MAX)"`
	Certifications string `json:"certifications" gorm:"column:certifications;type:varchar(MAX)"`

	CreatedAt      time.Time `json:"createdat" gorm:"column:createdat;autoCreateTime"`
	UpdatedAt      time.Time `json:"updatedat" gorm:"column:updatedat;autoUpdateTime"`
	Last_conection time.Time `json:"last_conection" gorm:"column:last_conection;type:datetime"`
	TokenSt        string    `json:"tokenst" gorm:"column:tokenst;type:varchar(255)"`
	Gym_ID         uuid.UUID `json:"gym_id" gorm:"column:gym_id;type:uniqueidentifier"`

	Users []USER `gorm:"foreignKey:coach_id"`
}

func (COACH) TableName() string {
	return "COACHES"
}

func (u COACH) updateLastConection() error {
	db, err := Get_db(false, false)
	if err != nil {
		return fmt.Errorf("error getting the database at: update_last_conection")
	}
	return db.Model(&COACH{}).Where("id = ?", u.ID).Update("last_connection", time.Now()).Error
}

func (u COACH) verifyHashedPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

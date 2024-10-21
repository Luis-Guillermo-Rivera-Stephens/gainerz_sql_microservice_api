package data

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type GYM struct {
	ID           uuid.UUID  `json:"id"  gorm:"column:id;type:uniqueidentifier;default:NEWID();primaryKey"`
	Distictive   string     `json:"name" gorm:"column:distinctive;type:varchar(255);not null;"`
	Email        string     `json:"email" gorm:"column:email;type:varchar(100)"`
	Password     string     `json:"password" gorm:"column:password;type:varchar(255)"`
	CreatedAt    time.Time  `json:"createdat" gorm:"column:createdat;autoCreateTime"`
	Latitud      float32    `json:"latitud" gorm:"column:latitud;float"`
	Longitud     float32    `json:"longitud" gorm:"column:longitud;float"`
	Address      string     `json:"address" gorm:"column:address;type:varchar(300)"`
	Address_link string     `json:"address_link" gorm:"column:address_link;type:varchar(300)"`
	Phone        string     `json:"phone" gorm:"column:phone;type:varchar(20)"`
	GymChain_ID  uuid.UUID  `json:"gymchain_id" gorm:"column:gymchain_id;type:uuid"`
	Users        []USER     `gorm:"foreignKey:gym_id"`
	Coaches      []COACH    `gorm:"foreignKey:gym_id"`
	GymImages    []GYMIMAGE `gorm:"foreignKey:gym_id"`
}

func (GYM) TableName() string {
	return "GYMS"
}

func (g GYM) verifyHashedPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(g.Password), []byte(password))
}

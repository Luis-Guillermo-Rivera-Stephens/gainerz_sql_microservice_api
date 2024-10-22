package data

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type USER struct {
	ID                          uuid.UUID `json:"id"  gorm:"column:id;type:uniqueidentifier;default:NEWID();primaryKey"`
	Name                        string    `json:"name" gorm:"column:name;type:varchar(255);not null;"`
	Email                       string    `json:"email" gorm:"column:email;type:varchar(255);uniqueIndex"`
	Password                    string    `json:"password" gorm:"column:password;type:varchar(255);not null"`
	Gender                      string    `json:"gender" gorm:"column:gender;type:varchar(50);not null"`
	Age                         int32     `json:"age" gorm:"column:age;type:int;not null"`
	Height                      int32     `json:"height" gorm:"column:height;type:int;not null"`
	Weight                      float32   `json:"weight" gorm:"column:weight;type:float;not null"`
	Injuries_and_contradictions string    `json:"injuries_and_contradictions" gorm:"column:injuries_and_contradictions;type:varchar(MAX)"`
	Objective                   string    `json:"objective" gorm:"column:objective;type:varchar(MAX)"`
	CreatedAt                   time.Time `json:"createdat" gorm:"column:createdat;autoCreateTime"`
	UpdatedAt                   time.Time `json:"updatedat" gorm:"column:updatedat;autoUpdateTime"`
	Last_conection              time.Time `json:"last_conection" gorm:"column:last_conection;type:datetime"`
	Status                      bool      `json:"status" gorm:"column:status;type:bit"`
	TokenSt                     string    `json:"tokenst" gorm:"column:tokenst;type:varchar(255)"`
	Coach_ID                    uuid.UUID `json:"coach_id" gorm:"column:coach_id;type:uniqueidentifier"`
	Gym_ID                      uuid.UUID `json:"gym_id" gorm:"column:gym_id;type:uniqueidentifier"`
	VIPS                        []VIP     `gorm:"foreignKey:user_id"`
}

func (USER) TableName() string {
	return "USERS"
}

func (u USER) updateLastConection() error {
	db, err := Get_db(Of_Flag)
	if err != nil {
		return fmt.Errorf("error getting the database at: update_last_conection")
	}
	return db.Model(&USER{}).Where("id = ?", u.ID).Update("last_connection", time.Now()).Error
}

func (u USER) verifyHashedPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

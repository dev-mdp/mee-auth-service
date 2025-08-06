package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	FirstName       string    `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName        string    `gorm:"type:varchar(255);not null" json:"last_name"`
	Email           string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	Password        string    `gorm:"type:varchar(255);not null" json:"password"`
	Address         string    `gorm:"type:varchar(255);not null" json:"address"`
	Address2        string    `gorm:"type:varchar(255);not null" json:"address2"`
	ProfilePath     string    `gorm:"type:varchar(255)" json:"profile_path"`
	Role            string    `gorm:"type:char(36);not null;foreginkey" json:"role"`
	Jtoken          string    `gorm:"type:varchar(255);not null" json:"jtoken"`
	EmailVerifiedAt time.Time `gorm:"type:timestamp" json:"email_verified_at"`
	Status          int16     `gorm:"default:0; type:smallint" json:"status"`
	CreatedAt       time.Time // otomatis dibuat GORM
	UpdatedAt       time.Time // otomatis dibuat GORM
}

func (user *User) BeforeSave(tx *gorm.DB) error {
	user.ID = uuid.New()

	return nil
}

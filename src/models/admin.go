package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type UserAdmin struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	FirstName       string    `gorm:"type:varchar(255);not null" json:"first_name"`
	LastName        string    `gorm:"type:varchar(255);not null" json:"last_name"`
	Email           string    `gorm:"type:varchar(255);not null;unique" json:"email"`
	PhoneNumber     string    `gorm:"type:varchar(255);not null" json:"phone_number"`
	Password        string    `gorm:"type:varchar(255);not null" json:"password"`
	ProfilePath     string    `gorm:"type:varchar(255)" json:"profile_path"`
	Jtoken          string    `gorm:"type:varchar(255);not null" json:"jtoken"`
	EmailVerifiedAt time.Time `gorm:"type:timestamp" json:"email_verified_at"`
	Status          int16     `gorm:"default:0; type:smallint" json:"status"`
	CreatedAt       time.Time // otomatis dibuat GORM
	UpdatedAt       time.Time // otomatis dibuat GORM
}

func (admin *UserAdmin) BeforeSave(tx *gorm.DB) error {
	admin.ID = uuid.New()

	return nil
}

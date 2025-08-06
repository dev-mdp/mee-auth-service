package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type MasterStatus struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Code      string    `gorm:"type:char(36); not null" json:"code"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	CreatedAt time.Time // otomatis dibuat GORM
	UpdatedAt time.Time // otomatis dibuat GORM
}

func (status *MasterStatus) BeforeSave(tx *gorm.DB) error {
	status.ID = uuid.New()

	return nil
}

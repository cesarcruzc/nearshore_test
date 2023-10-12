package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Device struct {
	ID        string     `json:"id" gorm:"type:char(36);primary_key;not null;unique_index"`
	Name      string     `json:"name" gorm:"type:varchar(100);not null"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

func (d *Device) BeforeCreate(tx *gorm.DB) (err error) {

	if d.ID == "" {
		d.ID = uuid.New().String()
	}

	return
}

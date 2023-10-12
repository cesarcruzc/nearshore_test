package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Firmware struct {
	ID           string     `json:"id" gorm:"type:char(36);primary_key;not null;unique_index"`
	Name         string     `json:"name" gorm:"type:varchar(100);not null"`
	DeviceID     string     `json:"device_id,omitempty" gorm:"type:char(36);not null;index"`
	Version      string     `json:"version" gorm:"type:varchar(10)"`
	ReleaseNotes string     `json:"release_notes" gorm:"type:text"`
	ReleaseDate  time.Time  `json:"release_date" gorm:"type:date"`
	Url          string     `json:"url" gorm:"type:varchar(100)"`
	CreatedAt    *time.Time `json:"-"`
	UpdatedAt    *time.Time `json:"-"`
}

func (f *Firmware) BeforeCreate(tx *gorm.DB) (err error) {

	if f.ID == "" {
		f.ID = uuid.New().String()
	}

	return
}

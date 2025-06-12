package domain

import "time"

type Image struct {
	ID  uint   `gorm:"primary_key"`
	URL string `gorm:"type:text;unique;not null;"`

	TodoID *uint `gorm:"index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time
}

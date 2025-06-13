package domain

import "time"

type Todo struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"type:text;size:100;unique;not null;"`
	Description string `gorm:"type:text;size:500;not null;"`
	Done        bool   `gorm:"type:boolean;not null;"`

	CategoryID uint     `gorm:"not null;index"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time
}

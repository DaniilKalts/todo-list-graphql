package domain

import "time"

type Category struct {
	ID          uint   `gorm:"primary_key"`
	Name        string `gorm:"type:text;size:100;not null;unique;"`
	Description string `gorm:"type:text;size:500;not null;"`

	Todos []Todo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	DeletedAt time.Time
}

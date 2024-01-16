package models

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	Title       string   `gorm:"not null;type:varchar(100)"`
	Description string   `gorm:"not null;type:varchar(1000)"`
	File        []string `gorm:"type:varchar[]"`
	SubItems    []SubItem
}

package models

import "gorm.io/gorm"

type SubItem struct {
	gorm.Model
	Description string `gorm:"not null;type:varchar(1000)"`
	ItemId      uint
}

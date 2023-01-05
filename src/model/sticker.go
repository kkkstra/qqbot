package model

import (
	"gorm.io/gorm"
)

type Sticker struct {
	gorm.Model
	Tag  string `gorm:"primarykey;not null"`
	File string `gorm:"not null"`
}

type StickerModel struct {
	db *gorm.DB
}

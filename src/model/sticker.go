package model

import (
	"gorm.io/gorm"
)

type Sticker struct {
	gorm.Model
	Tag  string `gorm:"primarykey;not null"`
	File string `gorm:"not null"`
	Url  string `gorm:"not null"`
}

type StickerModel struct {
	db *gorm.DB
}

func (m *StickerModel) CreateSticker(s *Sticker) error {
	res := m.db.Create(s)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (m *StickerModel) DeleteSticker(file string) error {
	tx := db.Where("file = ?", file).Delete(&Sticker{})
	return tx.Error
}

func (m *StickerModel) GetStickerByTag(tag string) (*[]Sticker, error) {
	res := &[]Sticker{}
	tx := db.Where("tag = ?", tag).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

func (m *StickerModel) GetStickerByFile(file string) (*[]Sticker, error) {
	res := &[]Sticker{}
	tx := db.Where("file = ?", file).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return res, nil
}

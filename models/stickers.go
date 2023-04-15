package models

import (
	"gorm.io/gorm"
)

type Stickers struct {
	gorm.Model
	Name string `gorm:"index"`
	URL  string
}

func AddSticker(db *gorm.DB, name, url string) (Stickers, error) {
	sticker := Stickers{Name: name, URL: url}

	result := db.Create(&sticker)
	if result.Error != nil {
		return Stickers{}, result.Error
	}

	return sticker, nil
}

func FindStickerByName(db *gorm.DB, name string) (Stickers, error) {
	var sticker Stickers
	result := db.Where("name = ?", name).First(&sticker)
	if result.Error != nil {
		return Stickers{}, result.Error
	}

	return sticker, nil
}

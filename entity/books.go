package entity

import "gorm.io/gorm"

type Books struct {
	gorm.Model
	Name_book    string
	Content_book string
	User_id      uint
}

type AksesBooks struct {
	DB *gorm.DB
}

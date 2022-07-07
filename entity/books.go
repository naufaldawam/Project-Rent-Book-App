package entity

import (
	"gorm.io/gorm"
	"log"
)

type Books struct {
	gorm.Model
	Name_book string
	User_id uint
}

type Pemilik struct {
	gorm.Model
	
}

type AksesBooks struct {
	DB *gorm.DB
}

// Function create book.
// =====================
func (ab *AksesBooks) NewBook(newBook Books) Books {
	err := ab.DB.Create(&newBook).Error
	if err != nil {
		log.Println(err)
		return Books{}
	}

	return newBook
}

// Function kepemilikan buku
// =====================
// func (ab *AksesBooks) PemilikBuku(newBook Books) bool{
// 	err := au.DB.Where("email = ? || phone = ?", Email, Phone).First(&Users{})

// }
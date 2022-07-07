package entity

import (
	"log"

	"gorm.io/gorm"
)

type Books struct {
	gorm.Model
	User_id    uint
	Name_book  string
	Rent_Books []Rent_Books `gorm:"foreignKey:Book_id"`
}

// select * from books dari go ke db:
// ==================================
type AksesBooks struct {
	DB *gorm.DB
}

// Function create book.
// =====================
func (ab *AksesBooks) AddBook(addBook Books) Books {
	err := ab.DB.Create(&addBook).Error
	if err != nil {
		log.Println(err)
		return Books{}
	}

	return addBook
}

package entity

import (
	"time"

	"gorm.io/gorm"
)

type Rent_Books struct {
	ID          int `gorm:"primaryKey"`
	Book_id     uint
	User_id     uint
	Rent_date   time.Time `gorm:"autoCreateTime"`
	Return_date time.Time
}

// select * from rent_books dari go ke db:
// =======================================
type AksesRentBooks struct {
	DB *gorm.DB
}

package config

import (
	"Project_Group/entity"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/rent_book_app?charset=utf8mb4&parseTime=True"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db

	// buat koneksi ke DB
	// <username>:<password>@tcp(<hostname>:<portDB>)/<db_name>
	// dbConnection := os.Getenv("DB_CONNECTION_RENT_BOOK_APP")
	// db, err := gorm.Open(mysql.Open(dbConnection), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal(err)
	// 	return nil
	// }
	// return db
}

func MigrateDB(conn *gorm.DB) {
	conn.AutoMigrate(entity.Users{})
}

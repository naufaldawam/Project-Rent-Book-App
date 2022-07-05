package entity

import (
	"log"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name_user string
	Email     string
	Phone     string
	Pass      string
}


//select * from user dari go ke db
type AksesUsers struct {
	DB *gorm.DB
}


// function create users
//=======================
func (au *AksesUsers) RegisterUser(newUser Users) Users {
	err := au.DB.Create(&newUser).Error
	if err != nil {
		log.Println(err)
		return Users{}
	}

	return newUser
}


//function validasi login
//=======================
func (au *AksesUsers) Validation(wasCreated Users) Users {
	err := au.DB.Select(&wasCreated)
}

//function read users
//=======================
// func (au *AksesUsers) GetAllData() []Users {
// 	var getUsers = []Users{}
// 	err := au.DB.Find(&getUsers)
// 	if err.Error != nil {
// 		log.Fatal(err.Statement.SQL.String())
// 		return nil
// 	}

// 	return getUsers
// }
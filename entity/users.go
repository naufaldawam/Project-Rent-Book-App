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
	Books     []Books `gorm:"foreignKey:User_id"`
}

// select * from user dari go ke db:
type AksesUsers struct {
	DB *gorm.DB
}

// Function create user.
// =====================
func (au *AksesUsers) RegisterUser(newUser Users) Users {
	err := au.DB.Create(&newUser).Error
	if err != nil {
		log.Println(err)
		return Users{}
	}

	return newUser
}

//function validasi create user
func (au *AksesUsers) IsCreated(Email, Phone string) bool {
	err := au.DB.Where("email = ? OR phone = ?", Email, Phone).First(&Users{})
	if err.Error != nil && err.Error != gorm.ErrRecordNotFound {
		return true
	}
	if err.RowsAffected != 0 {
		return true
	}
	return false
}

// Function auth email.
// ====================
func (au *AksesUsers) CekEmail(Email string) bool {
	cekEmail := au.DB.Where("email = ?", Email).Find(&Users{})
	if err := cekEmail.Error; err != nil {
		return true
	}
	if aff := cekEmail.RowsAffected; aff != 0 {
		return true
	}
	return false
}

// Function auth password.
// ====================
func (au *AksesUsers) CekPassword(Pass string) bool {
	cekPass := au.DB.Where("pass = ?", Pass).Find(&Users{})
	if err := cekPass.Error; err != nil {
		return true
	}
	if aff := cekPass.RowsAffected; aff != 0 {
		return true
	}
	return false
}

//function read users
//===================
func (au *AksesUsers) GetAllUsers(Name string) Users {
	var GetAllUsers = Users{}
	err := au.DB.Select("name", "email", "phone", "created_at").Find(&GetAllUsers)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
	}
	return GetAllUsers
}

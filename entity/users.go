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
	Books   []Books `gorm:"foreignKey:User_id"`
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
	err := au.DB.Where("email = ? || phone = ?", Email, Phone).First(&Users{})
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
	if err := cekEmail.Error ; err != nil{
		return false
	}
	if aff := cekEmail.RowsAffected; aff != 0 {
		return false
	}
	return true
}


// Function auth email.
// ====================
func (au *AksesUsers) CekPassword(Password string) bool {
	cekPass := au.DB.Where("email = ?", Password).Find(&Users{})
	if err := cekPass.Error ; err != nil{
		return false
	}
	if aff := cekPass.RowsAffected; aff != 0 {
		return false
	}
	return true
}

//function read users
//===================
func (au *AksesUsers) GetAllData() []Users {
	var getUsers = []Users{}
	err := au.DB.Find(&getUsers)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return getUsers
}

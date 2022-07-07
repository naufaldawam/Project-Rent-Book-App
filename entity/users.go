package entity

import (
	"log"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name_user  string
	Email      string
	Phone      string
	Pass       string
	Books      []Books      `gorm:"foreignKey:User_id"`
	Rent_Books []Rent_Books `gorm:"foreignKey:User_id"`
}

// select * from users dari go ke db:
// ==================================
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

// Function validasi create user.
// ==============================
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

// Function login user.
// ====================
func (au *AksesUsers) Islogin(Email, Pass string) bool {
	err := au.DB.Where("email = ? AND pass = ?", Email, Pass).First(&Users{})
	if err.Error != nil {
		return true
	}
	return false
}

func (au *AksesUsers) DeleteUser(Id int) bool {
	postExc := au.DB.Where("id = ?", Id).Delete(&Users{})
	if err := postExc.Error; err != nil {
		log.Fatal(err)
		return false
	}
	if aff := postExc.RowsAffected; aff < 1 {
		log.Println("Tidak ada data yang dihapus")
		return false
	}
	return true
}

// Function read users.
// ====================
func (au *AksesUsers) GetAllData() []Users {
	var getUsers = []Users{}
	err := au.DB.Find(&getUsers)
	if err.Error != nil {
		log.Fatal(err.Statement.SQL.String())
		return nil
	}

	return getUsers
}

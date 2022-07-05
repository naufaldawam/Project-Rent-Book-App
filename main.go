package main

import (
	"Project_Group/config"
	"Project_Group/entity"
	"fmt"
	"os/user"
)



func main() {
	conn := config.InitDB()
	aksesUsers := entity.AksesUsers{DB: conn}
	var input int = 0

	for input != 4 {
		
		fmt.Println("\t Rent Book App")
		fmt.Println("1. Login App")
		fmt.Println("2. Register")
		fmt.Println("3. Katalog Buku")
		fmt.Println("4. Keluar")
		fmt.Print("Masukkan Pilihan Menu: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			// var login entity.Users

		case 2:
			var newUser entity.Users
			// var input string
			fmt.Print("Masukkan username: ")
			fmt.Scanln(&newUser.Name_user)
			fmt.Print("Masukkan email: ")
			fmt.Scanln(&newUser.Email)
			fmt.Print("Masukkan nomor hp: ")
			fmt.Scanln(&newUser.Phone)
			fmt.Print("Masukkan password: ")
			fmt.Scanln(&newUser.Pass)
			
			
		default:
			continue
		}
	}
	fmt.Println("Program telah berhenti!")
}

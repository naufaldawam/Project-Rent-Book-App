package main

import (
	"Project_Group/config"
	"Project_Group/entity"
	"fmt"
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
			var login entity.Users
			fmt.Print("Masukkan Username: ")
			fmt.Scanln(&login.Email)
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&login.Pass)
			isLogin := aksesUsers.Islogin(login.Email, login.Pass)
			if !isLogin {
				fmt.Println("Anda berhasil Login!")
			} else {
				fmt.Println("Anda gagal Login!")
			}

		case 2:
			var newUser entity.Users
			fmt.Print("Masukkan Username: ")
			fmt.Scanln(&newUser.Name_user)
			fmt.Print("Masukkan Email: ")
			fmt.Scanln(&newUser.Email)
			fmt.Print("Masukkan Nomor HP: ")
			fmt.Scanln(&newUser.Phone)
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&newUser.Pass)
			var isCreated = aksesUsers.IsCreated(newUser.Email, newUser.Phone)
			if isCreated == false {
				res := aksesUsers.RegisterUser(newUser)
				fmt.Println("Selamat bergabung!", res.Name_user)
				break
			} else {
				fmt.Println("Tidak bisa input, karena Error!")
			}

		default:
			continue
		}
	}
	fmt.Println("Program telah berhenti!")
}

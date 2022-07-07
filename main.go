package main

import (
	"Project_Group/config"
	"Project_Group/entity"
	"fmt"
	// "os/user"
)

func login() (Email string, Pass string) {
	fmt.Println("Anda memilih menu login , Silahkan login terlebih dahulu")
	fmt.Println("Masukan Email: ")
	fmt.Scanln(&Email)
	fmt.Println("Masukan Password")
	fmt.Scanln(&Pass)
	return Email, Pass
}

func main() {
	conn := config.InitDB()
	aksesUsers := entity.AksesUsers{DB: conn}
	// aksesBooks := entity.AksesBooks{DB: conn}
	var menu = false
	for !menu {
		var input int = 0
		fmt.Println("============================")
		fmt.Println("\t Rent Book App")
		fmt.Println("1. Login App")
		fmt.Println("2. Register")
		fmt.Println("3. Katalog Buku")
		fmt.Println("4. Keluar")
		fmt.Print("Masukkan Pilihan Menu: ")
		fmt.Scanln(&input)

		if input == 1 {
			Email, Pass := login()
			emailAuth := aksesUsers.CekEmail(Email)
			passAuth := aksesUsers.CekPassword(Pass)

			if !emailAuth && !passAuth {
				fmt.Println("Akun tidak ditemukan \n silahkan register terlebih dahulu")
			} else if !emailAuth || !passAuth {
				fmt.Println("Email dan password salah")
			} else {
				fmt.Println("Anda berhasil login")
				menu = true
			}

		} else if input == 2 {
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
			if !isCreated {
				res := aksesUsers.RegisterUser(newUser)
				fmt.Println("Selamat bergabung!", res.Name_user)
			} else {
				fmt.Println("Tidak bisa input, karena Error!")
			}
		} else if input == 3 {
			fmt.Println("tampil buku")
		} else {
			fmt.Println("terima kasih sudah menggunakan aplikasi ini")
			break
		}

	}
	// for menu {
	// 		// fmt.Println("ini menu")
	// 	}

	// conn := config.InitDB()
	// aksesUsers := entity.AksesUsers{DB: conn}
	// aksesBooks := entity.AksesBooks{DB: conn}
	// var input int = 0

	// for input != 4 {
	// 	fmt.Println("\t Rent Book App")
	// 	fmt.Println("1. Login App")
	// 	fmt.Println("2. Register")
	// 	fmt.Println("3. Katalog Buku")
	// 	fmt.Println("4. Keluar")
	// 	fmt.Print("Masukkan Pilihan Menu: ")
	// 	fmt.Scanln(&input)

	// 	switch input {
	// 	case 1:
	// 		var login entity.Users
	// 		fmt.Print("Masukkan Username: ")
	// 		fmt.Scanln(&login.Email)
	// 		fmt.Print("Masukkan Password: ")
	// 		fmt.Scanln(&login.Pass)
	// 		isLogin := aksesUsers.Islogin(login.Email, login.Pass)
	// 		if !isLogin {
	// 			fmt.Println("Anda berhasil Login!")
	// 		} else {
	// 			fmt.Println("Anda gagal Login!")
	// 		}

	// 	case 2:
	// 		var newUser entity.Users
	// 		fmt.Print("Masukkan Username: ")
	// 		fmt.Scanln(&newUser.Name_user)
	// 		fmt.Print("Masukkan Email: ")
	// 		fmt.Scanln(&newUser.Email)
	// 		fmt.Print("Masukkan Nomor HP: ")
	// 		fmt.Scanln(&newUser.Phone)
	// 		fmt.Print("Masukkan Password: ")
	// 		fmt.Scanln(&newUser.Pass)
	// 		var isCreated = aksesUsers.IsCreated(newUser.Email, newUser.Phone)
	// 		if !isCreated {
	// 			res := aksesUsers.RegisterUser(newUser)
	// 			fmt.Println("Selamat bergabung!", res.Name_user)
	// 			break
	// 		} else {
	// 			fmt.Println("Tidak bisa input, karena Error!")
	// 		}
	// 		break

	// 	case 3:
	// 		var newBook entity.Books
	// 		fmt.Println("Masukan Buku: ")
	// 		fmt.Scanln(&newBook.Name_book)
	// 		fmt.Println("Masukan kepemilikan buku: ")
	// 		fmt.Scanln(&newBook.User_id)
	// 		res := aksesBooks.NewBook(newBook)
	// 		fmt.Println("Buku berhasil di input: ", res.Name_book)
	// 		break

	// 	case 4:
	// 		break

	// 	default:
	// 		fmt.Println("===============================\t")
	// 		fmt.Println("harap masukan menu yang sesuai!")
	// 		fmt.Println("===============================\t")
	// 		continue
	// 	}
	// }
	// fmt.Println("Program telah berhenti!")
}

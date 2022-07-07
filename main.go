package main

import (
	"Project_Group/config"
	"Project_Group/entity"
	"fmt"
	// "os/user"
)

func login() (string, string) {
	var email, pass string
	fmt.Println("Anda memilih menu login , Silahkan login terlebih dahulu")
	fmt.Println("Masukan Email: ")
	fmt.Scanln(&email)
	fmt.Println("Masukan Password")
	fmt.Scanln(&pass)
	return email, pass
}

func main() {
	conn := config.InitDB()
	config.MigrateDB(conn)
	aksesUsers := entity.AksesUsers{DB: conn}
	// aksesBooks := entity.AksesBooks{DB: conn}
	// var userActive aksesUsers.Users

	var menu = false
	for !menu {
		var input int = 0
		fmt.Println("============================")
		fmt.Println("\t Menu Awal Rent Book App")
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
			fmt.Println("email: ", Email)
			fmt.Println("pass: ", Pass)
			fmt.Println("auth email: ", emailAuth)
			fmt.Println("auth pass: ", passAuth)

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
			fmt.Println("is created : ", isCreated)
			if isCreated == true {
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
	for menu {
		var input int
		fmt.Println("--------------- SELAMAT DATANG ---------------")
		fmt.Println("----- Silahkan Pilih Fitur yang Tersedia -----")
		fmt.Println("1. Lihat Akun Saya")
		fmt.Println("2. Perbarui Akun Saya")
		fmt.Println("3. Hapus Akun Saya")
		// fmt.Println("4. Tambah Buku Saya")
		// fmt.Println("5. Lihat Daftar Buku Anda")
		// fmt.Println("6. Perbarui Buku Anda")
		// fmt.Println("7. Hapus Buku Anda")
		// fmt.Println("8. Pinjam Buku")
		// fmt.Println("9. Kembalikan Buku")
		// fmt.Println("10. Lihat Daftar Buku Yang Tersedia")
		fmt.Println("11. Logout")
		fmt.Print("Pilih Menu: ")
		fmt.Scan(&input)
		fmt.Print("\n")

		switch input {
		case 1: //liat Akun
			fmt.Println("----- Info Akun Saya -----")
		// for _, val := range aksesUsers.GetAllUsers(){
		// 	fmt.Println("Nama: ", val.Name)
		// 	fmt.Println("Nomor HP: ", val.Phone)
		// 	fmt.Println("User Name: ", val.Name)
		// 	fmt.Println("Email: ", val.Email)
		// }
		case 2:
			fmt.Println("----- Perbaharui Akun Saya -----")


		case 11:
			menu = false
		}

	}

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

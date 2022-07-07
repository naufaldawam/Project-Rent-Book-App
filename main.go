package main

import (
	"Project_Group/config"
	"Project_Group/entity"
	"fmt"
)

func main() {
	conn := config.InitDB()
	aksesUsers := entity.AksesUsers{DB: conn}
	aksesBooks := entity.AksesBooks{DB: conn}
	var menu int
	for menu != 4 {
		fmt.Println("\n\t 📌 Rent Book App 📌")
		fmt.Println("\t ===================")
		fmt.Println("1️⃣  Login App")
		fmt.Println("2️⃣  Registrasi User")
		fmt.Println("3️⃣  Katalog Buku")
		fmt.Println("4️⃣  Keluar")
		fmt.Println("======================")
		fmt.Print("Masukkan pilihan menu: ")
		fmt.Scanln(&menu)
		// Menu awal Rent Book App:
		switch menu {
		case 1:
			var login entity.Users
			fmt.Println("======================")
			fmt.Print("Masukkan Username: ")
			fmt.Scanln(&login.Email)
			fmt.Print("Masukkan Password: ")
			fmt.Scanln(&login.Pass)
			isLogin := aksesUsers.Islogin(login.Email, login.Pass)
			for !isLogin {
				var menu2 int
				fmt.Println("\n\t📌 Selamat Datang 📌")
				fmt.Println("\t ===================")
				fmt.Println("1️⃣  Nonaktifkan User")
				fmt.Println("2️⃣  Tambah Buku")
				fmt.Println("3️⃣  Update Buku")
				fmt.Println("4️⃣  Pinjam Buku")
				fmt.Println("9️⃣ 9️⃣  Keluar")
				fmt.Println("======================")
				fmt.Print("Masukkan pilihan menu: ")
				fmt.Scanln(&menu2)
				// Menu ketika sudah login:
				switch menu2 {
				case 1:
					var ID int
					fmt.Println("======================")
					fmt.Print("Masukkan ID User: ")
					fmt.Scanln(&ID)
					fmt.Println(aksesUsers.DeactiveUser(ID))

				case 99:
					isLogin = true

				default:
					continue
				}
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
			if !isCreated {
				res := aksesUsers.RegisterUser(newUser)
				fmt.Println("Selamat bergabung!", res.Name_user)
			} else {
				fmt.Println("Tidak bisa input, karena Error!")
			}

		case 3:
			var addBook entity.Books
			fmt.Print("Masukan Judul Buku: ")
			fmt.Scanln(&addBook.Name_book)
			fmt.Print("Masukan Kepemilikan Buku: ")
			fmt.Scanln(&addBook.User_id)
			res := aksesBooks.AddBook(addBook)
			fmt.Println("Buku berhasil di input: ", res.Name_book)

		case 4:
			break

		default:
			fmt.Println("===============================\t")
			fmt.Println("Harap masukan menu yang sesuai!")
			fmt.Println("===============================\t")
			continue
		}
	}
	fmt.Println("======================")
	fmt.Println("⚠️  Program telah berhenti!⚠️")
}

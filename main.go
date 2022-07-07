package main

import (
	"Project_Group/config"
	"Project_Group/entity"
	"bufio"
	"fmt"
	"os"
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
	var userActive entity.Users

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
			// fmt.Println("is created : ", isCreated)
			if isCreated {
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
		fmt.Println("00. Logout")
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
			var IDUSER = userActive.Id_user

			updateUser := entity.Users{}

			fmt.Println("----- Perbaharui Akun Saya -----")
			fmt.Print("Masukkan Nama: ")
			Name_user := bufio.NewReader(os.Stdin)
			updateUser.Name_user, _ = Name_user.ReadString('\n')
			fmt.Print("Masukkan Email: ")
			Email := bufio.NewReader(os.Stdin)
			updateUser.Email, _ = Email.ReadString('\n')
			fmt.Print("Masukkan Password: ")
			Pass := bufio.NewReader(os.Stdin)
			updateUser.Pass, _ = Pass.ReadString('\n')
			fmt.Print("Masukkan Nomor HP: ")
			Phone := bufio.NewReader(os.Stdin)
			updateUser.Phone, _ = Phone.ReadString('\n')

			// aksesUsers.EditUser(IDUSER, updateUser)
			// fmt.Println("Update Data Berhasil")


		case 00:
			menu = false
		}

	}
}

package main

import "fmt"

const NMAX int = 999

//struct untuk menyimpan data menu
type menu struct {
	id, namaMenu, kategori, status, komposisi string
	harga                                     float64
}
type menus [NMAX]menu

//fungsi untuk mencetak data menu
func cetakData(data menus, n int) {
	var i int
	fmt.Println("==========================================================================================================================")
	fmt.Println("|                                                  Daftar Menu                                                           |")
	fmt.Println("==========================================================================================================================")
	fmt.Printf("%-5s %-25s %-12s %-15s %-35s %-10s\n", "ID", "Nama Menu", "Kategori", "Status", "Komposisi", "Harga")
	fmt.Println("--------------------------------------------------------------------------------------------------------------------------")
	for i = 0; i < n; i++ {
		fmt.Printf("%-5s %-25s %-12s %-15s %-35s %.2f\n", data[i].id, data[i].namaMenu, data[i].kategori, data[i].status, data[i].komposisi, data[i].harga)
	}
}

//fungsi untuk mengurutkan data menu berdasarkan harga
func selectionSortHarga(data *menus, n int) {
	var i, j, min int
	for i = 0; i < n-1; i++ {
		min = i
		for j = i + 1; j < n; j++ {
			if data[j].harga < data[min].harga {
				min = j
			}
		}
		data[i], data[min] = data[min], data[i]
	}
}

//fungsi untuk mengurutkan data menu berdasarkan kategori dan harga dengan cara selection sort
func sortMenuByKategori(data *menus, n int) {
	var coffee, noncoffee menus
	var nCoffee, nNon int
	var i int

	nCoffee = 0
	nNon = 0
	for i = 0; i < n; i++ {
		if data[i].kategori == "coffee" {
			coffee[nCoffee] = data[i]
			nCoffee = nCoffee + 1
		} else if data[i].kategori == "noncoffee" {
			noncoffee[nNon] = data[i]
			nNon = nNon + 1
		}
	}
	insertionSortHarga(&coffee, nCoffee)
	selectionSortHarga(&noncoffee, nNon)
	for i = 0; i < nCoffee; i++ {
		data[i] = coffee[i]
	}
	for i = 0; i < nNon; i++ {
		data[nCoffee+i] = noncoffee[i]
	}
}

//fungsi untuk mengurutkan data menu berdasarkan harga dengan cara insertion sort
func insertionSortHarga(data *menus, n int) {
	var pass, i int
	var temp menu

	pass = 1
	for pass < n {
		i = pass
		temp = data[pass]
		for i > 0 && data[i-1].harga > temp.harga {
			data[i] = data[i-1]
			i = i - 1
		}
		data[i] = temp
		pass = pass + 1
	}
}

//fungsi untuk menampilkan menu utama
func homeCafe(aktivitas int) int {
	fmt.Println("")
	fmt.Println("=== HOME TELKOM CAFE ===")
	fmt.Println("1.Tambahkan Menu")
	fmt.Println("2.Tampilkan Menu & pemesanan")
	fmt.Println("3.ubah menu")
	fmt.Println("4.Hapus menu")
	fmt.Println("5.Statistik penjualan")
	fmt.Println("6.Cari menu berdasarkan kategori menu")
	fmt.Println("7.Keluar")
	fmt.Println("masukkan nomor aktivitas yang anda inginkan")
	return aktivitas
}

//fungsi untuk mengubah data menu
func ubahMenu(data *menus, n int) {
	var idx int
	var idCari string

	fmt.Print("Masukkan ID menu yang ingin diubah: ")
	fmt.Scan(&idCari)
	idx = sequentialSearch(*data, n, idCari)
	if idx != -1 {
		fmt.Print("Nama menu baru: ")
		fmt.Scan(&data[idx].namaMenu)
		fmt.Print("Status baru: ")
		fmt.Scan(&data[idx].status)
		fmt.Print("Komposisi baru: ")
		fmt.Scan(&data[idx].komposisi)
		fmt.Print("Harga baru: ")
		fmt.Scan(&data[idx].harga)
		fmt.Println("Data berhasil diubah")
	} else {
		fmt.Println("Menu tidak ditemukan")
	}
	if idx != -1 {
		sortMenuByKategori(data, n)
	}
}

//fungsi untuk menghapus data menu
func hapusMenu(data *menus, n *int) {
	var idx, i int
	var idCari string

	fmt.Print("Masukkan ID menu yang ingin dihapus: ")
	fmt.Scan(&idCari)
	idx = sequentialSearch(*data, *n, idCari)
	if idx != -1 {
		for i = idx; i < *n-1; i++ {
			data[i] = data[i+1]
		}

		*n = *n - 1

		fmt.Println("Data berhasil dihapus")
	} else {
		fmt.Println("Data tidak ditemukan")
	}
}

//fungsi untuk menghitung harga rata-rata menu dan banyak jenis minuman
func hargaRataMenu(data *menus, n int) {
	var hargaRata float64
	var i int
	hargaRata = 0

	for i = 0; i < n; i++ {
		hargaRata = hargaRata + data[i].harga
	}
	hargaRata = hargaRata / float64(n)
	fmt.Printf("Harga rata-rata menu: %.2f\n", hargaRata)
}

//fungsi untuk menghitung banyak jenis minuman
func banyakJenisMinuman(data *menus, n int) {
	var i, nCoffee, nNon int
	nCoffee = 0
	nNon = 0
	for i = 0; i < n; i++ {
		if data[i].kategori == "coffee" {
			nCoffee = nCoffee + 1
		} else if data[i].kategori == "noncoffee" {
			nNon = nNon + 1
		}
	}
	fmt.Println("Banyak menu coffee:")
	fmt.Println(nCoffee)
	fmt.Println("Banyak menu noncoffee:")
	fmt.Println(nNon)
}

//fungsi untuk mencari menu berdasarkan id menu
func sequentialSearch(data menus, n int, cari string) int {
	var i int
	for i = 0; i < n; i++ {
		if data[i].id == cari {
			return i
		}
	}
	return -1
}

//fungsi untuk mencari menu berdasarkan kategori menu dengan binary search
func binarySearch(data menus, n int, cari string) int {
	var kiri, kanan, tengah int
	kiri = 0
	kanan = n - 1
	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if data[tengah].kategori == cari {
			return tengah
		} else if data[tengah].kategori < cari {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func main() {
	var data menus
	var n, aktivitas int
	var idMenu, tambahinPesanan, pesan string

	n = 0
	fmt.Println("")
	fmt.Println("Selamat datang di cafe Telkom University")
	homeCafe(aktivitas)
	fmt.Scan(&aktivitas)

	//loop untuk menjalankan aktivitas sampai user memilih keluar(selama aktivitas tidak sama dengan 6)
	for aktivitas != 7 {
		//aktivitas 1 untuk menambahkan menu
		if aktivitas == 1 {
			fmt.Println("Masukkan id menu(a=coffee, b=noncoffee, END=stop):")
			fmt.Scan(&idMenu)

			for idMenu != "END" {
				data[n].id = idMenu
				fmt.Println("Masukkan nama menu:")
				fmt.Scan(&data[n].namaMenu)
				if data[n].id[0] == 'a' {
					data[n].kategori = "coffee"
				} else if data[n].id[0] == 'b' {
					data[n].kategori = "noncoffee"
				}
				fmt.Println("Status menu (tersedia/tidaktersedia):")
				fmt.Scan(&data[n].status)
				fmt.Println("Masukkan komposisi dari menu yang anda tambahkan")
				fmt.Scan(&data[n].komposisi)
				fmt.Println("Harga menu:")
				fmt.Scan(&data[n].harga)
				n = n + 1
				fmt.Println("Masukkan id menu(a=coffee, b=noncoffee, END=stop):")
				fmt.Scan(&idMenu)
			}
			sortMenuByKategori(&data, n)
			cetakData(data, n)
			if idMenu == "END" {
				homeCafe(aktivitas)
				fmt.Scan(&aktivitas)
			}
			//aktivitas 2 untuk menampilkan menu dan pemesanan
		} else if aktivitas == 2 {
			if n > 0 {
				fmt.Println("\n=== DATA MENU ===")
				cetakData(data, n)
				fmt.Println("Apakah anda ingin memesan?(iya/tidak)")
				fmt.Scan(&pesan)
				if pesan == "tidak" {
					homeCafe(aktivitas)
					fmt.Scan(&aktivitas)
				} else if pesan == "iya" {
					var again string
					var total float64
					var daftarP [NMAX]string
					var hargaP [NMAX]float64
					var jumlahP, i int
					var found bool

					again = "iya"
					jumlahP = 0
					for again == "iya" {
						fmt.Println("Masukkan nama menu yang ingin dipesan:")
						fmt.Scan(&pesan)
						found = false
						for i = 0; i < n; i++ {
							if data[i].namaMenu == pesan {
								found = true
								if data[i].status == "tersedia" {
									fmt.Printf("Pesanan %s berhasil dibuat dengan harga %.2f\n", data[i].namaMenu, data[i].harga)
									total = total + data[i].harga
									daftarP[jumlahP] = data[i].namaMenu
									hargaP[jumlahP] = data[i].harga
									jumlahP = jumlahP + 1
									fmt.Println("total harga pesanan anda: ", total)
									fmt.Println("Apakah ingin tambah pesanan lagi?(iya/tidak)")
									fmt.Scan(&again)
								} else {
									fmt.Printf("Maaf, menu %s tidak tersedia saat ini.\n", data[i].namaMenu)
								}
							}
						}
						if !found {
							fmt.Printf("Maaf, menu %s tidak ditemukan.\n", pesan)
						}
					}
					fmt.Println("\n=== MENU PESANAN ANDA ===")
					for i = 0; i < jumlahP; i++ {
						fmt.Printf("%d. %s %.2f\n", i+1, daftarP[i], hargaP[i])
					}
					fmt.Printf("Total harga = %.2f\n", total)
					homeCafe(aktivitas)
					fmt.Scan(&aktivitas)
				}
			} else {
				fmt.Println("belum ada daftar menu nih,tambahin menu dulu dong (oke/tidak):")
				fmt.Scan(&tambahinPesanan)

				if tambahinPesanan == "tidak" {
					aktivitas = 7
				} else {
					homeCafe(aktivitas)
					fmt.Scan(&aktivitas)
				}
			}
			//aktivitas 3 untuk mengubah menu
		} else if aktivitas == 3 {
			if n > 0 {
				ubahMenu(&data, n)
				cetakData(data, n)
				homeCafe(aktivitas)
				fmt.Scan(&aktivitas)
			} else {
				fmt.Println("belum ada daftar menu nih,tambahin menu dulu dong (oke/tidak):")
				fmt.Scan(&tambahinPesanan)
				if tambahinPesanan == "tidak" {
					aktivitas = 7
				} else {
					aktivitas = 1
				}
			}
			//aktivitas 4 untuk menghapus menu
		} else if aktivitas == 4 {
			if n > 0 {
				hapusMenu(&data, &n)
				cetakData(data, n)
				homeCafe(aktivitas)
				fmt.Scan(&aktivitas)
			} else {
				fmt.Println("belum ada daftar menu nih,tambahin menu dulu dong (oke/tidak):")
				fmt.Scan(&tambahinPesanan)
				if tambahinPesanan == "tidak" {
					aktivitas = 7
				} else {
					aktivitas = 1
				}
			}
			//aktivitas 5 untuk menampilkan statistik penjualan
		} else if aktivitas == 5 {
			if n > 0 {
				fmt.Println("")
				hargaRataMenu(&data, n)
				banyakJenisMinuman(&data, n)

				homeCafe(aktivitas)
				fmt.Scan(&aktivitas)
			} else {
				fmt.Println("belum ada daftar menu nih sehingga tidak bisa menampilkan statistik,mau tambah menu? (oke/tidak):")
				fmt.Scan(&tambahinPesanan)
				if tambahinPesanan == "tidak" {
					aktivitas = 7
				}
			}
		} else if aktivitas == 6 {
			var kategoriCari string
			if n > 0 {
				fmt.Println("Cari menu berdasarkan kategori:")
				fmt.Scan(&kategoriCari)
				binarySearch(data, n, kategoriCari)
				homeCafe(aktivitas)
				fmt.Scan(&aktivitas)
			} else {
				fmt.Println("belum ada daftar menu nih sehingga tidak bisa mencari menu berdasarkan kategori,mau tambah menu? (oke/tidak):")
				fmt.Scan(&tambahinPesanan)
				if tambahinPesanan == "tidak" {
					aktivitas = 7
				}
			}
		} else {
			fmt.Println("Aktivitas tidak valid. Silakan pilih nomor aktivitas yang tersedia.")
			homeCafe(aktivitas)
			fmt.Scan(&aktivitas)
		}
	}
	//pesan keluar jika user memilih aktivitas 7
	if aktivitas == 7 {
		fmt.Println("Aktivitas Selesai Terimakasih")
	}
}

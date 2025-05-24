package main

import "fmt"

const NMAX int = 10

type pinjaman struct {
	id       string
	nama     string
	pinjaman int
	tenor    int
	bunga    float64
	tBunga   float64
	kredit   float64
}
type tabPinjaman [NMAX]pinjaman

var data tabPinjaman
var nData int

func main() {
	fmt.Println("⎛⎝ ≽ > ⩊ < ≼ ⎠⎞      ﷽     ⎛⎝ ≽ > ⩊ < ≼ ⎠⎞")
	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("┃   Selamat datang di Aplikasi Pinjaman  ┃")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Println("┃   Aplikasi ini mensimulasikan sistem   ┃")
	fmt.Println("┃     pinjaman dan kredit sederhana      ┃")
	fmt.Println("┃       Tekan ENTER untuk memulai        ┃")
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Scanln()
	menu()
}

func menu() {
	var pilih int

	for {
		fmt.Println(" 𓂃˖˳·˖ ִֶָ ⋆🌷͙⋆ ִֶָ˖·˳˖𓂃   ִֶָ 𓂃˖˳·˖ ִֶָ ⋆🌷͙⋆ ִֶָ˖·˳˖𓂃 ִֶָ")
		fmt.Println("🌹            PINJAMAN BANK            🌹")
		fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
		fmt.Println("┃ [1] Tambah Data Peminjam               ┃")
		fmt.Println("┃ [2] Ubah / Hapus Data Peminjam         ┃")
		fmt.Println("┃ [3] Urutkan Daftar Peminjam            ┃")
		fmt.Println("┃ [4] Hitung Bunga & Cicilan             ┃")
		fmt.Println("┃ [5] Cari Data Peminjam                 ┃")
		fmt.Println("┃ [6] Tampilkan Laporan                  ┃")
		fmt.Println("┃ [0] Exit                               ┃")
		fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
		fmt.Print("Pilih No ➝  ")

		fmt.Scan(&pilih)
		fmt.Println()

		switch pilih {
		case 1:
			tambahData(&data, &nData)
		case 2:
			if nData == 0 {
				fmt.Println("Belum Ada Data Peminjam. Silakan Tambahkan Terlebih Dahulu!")
			} else {
				pilihUbahHapusData()
			}
		case 3:
			if nData == 0 {
				fmt.Println("Belum Ada Data Untuk Diurutkan. Silakan Tambahkan Terlebih Dahulu!")
			} else {
				pilihSort()
			}
		case 4:
			if nData == 0 {
				fmt.Println("Belum Ada Data Untuk Dihitung. Silakan Tambahkan Terlebih Dahulu!")
			} else {
				hitungBunga(&data, nData)
			}
		case 5:
			if nData == 0 {
				fmt.Println("Belum Ada Data Untuk Dicari. Silakan Tambahkan Terlebih Dahulu!")
			} else {
				pilihCari()
			}
		case 6:
			if nData == 0 {
				fmt.Println("Belum Ada Data Untuk Ditampilkan. Silakan Tambahkan Terlebih Dahulu!")
			} else {
				cetakKredit(data, nData)
			}
		case 0:
			fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
			fmt.Println("┃    TERIMA KASIH! SAMPAI JUMPA LAGI!    ┃")
			fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
		default:
			fmt.Println("Pilihan Tidak Tersedia. Silakan Coba Lagi!")
		}
		if pilih == 0 {
			break
		}
	}
}

func tambahData(A *tabPinjaman, n *int) {
	var i, jumlah int
	var idP string

	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("┃         TAMBAH DATA PEMINJAM BARU         ┃")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Println("┃ PENJELASAN                                ┃")
	fmt.Println("┃ Setiap peminjam punya:                    ┃")
	fmt.Println("┃ ID unik, nama, jumlah pinjaman, dan tenor ┃")
	fmt.Println("┃ NOTE: Masukkan maksimal 5 huruf untuk ID  ┃")
	fmt.Println("┃       dan maksimal 2 kata untuk nama      ┃")
	fmt.Println("┃                                           ┃")
	fmt.Println("┃ Contoh                                    ┃")
	fmt.Println("┃  ID Unik                : A01             ┃")
	fmt.Println("┃  Nama Peminjam          : Deni_Saepudin   ┃")
	fmt.Println("┃  Jumlah Pinjaman (Rp)   : 10000000        ┃")
	fmt.Println("┃  Tenor Pinjaman (bulan) : 12              ┃")
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")

	fmt.Print("Jumlah data yang ingin dimasukkan (max 10): ")
	fmt.Scan(&jumlah)

	if *n < NMAX {
		if jumlah > NMAX {
			jumlah = NMAX
			fmt.Println("\033[31mKapasitas maksimal tercapai. Data yang ditambahkan dibatasi menjadi 10!\033[0m")
		}

		for i = 0; i < jumlah; i++ {
			fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
			fmt.Printf("Data peminjam ke-%d\n", *n+1)

			for {
				fmt.Print("ID unik                : ")
				fmt.Scan(&idP)
				if !IDSama(*A, *n, idP) {
					A[*n].id = idP
					break
				} else {
					fmt.Println("ID sudah digunakan, Silakan masukkan ID lain!")
				}
			}

			fmt.Print("Nama peminjam          : ")
			fmt.Scan(&A[*n].nama)

			fmt.Print("Jumlah pinjaman (Rp)   : ")
			fmt.Scan(&A[*n].pinjaman)

			fmt.Print("Tenor pinjaman (bulan) : ")
			fmt.Scan(&A[*n].tenor)

			*n++
		}
		fmt.Println("\nData berhasil ditambahkan!")
	} else {
		fmt.Println("Data tidak dapat ditambahkan, kapasitas sudah penuh!")
	}

}

func IDSama(A tabPinjaman, n int, idP string) bool {
	var i int
	for i = 0; i < n; i++ {
		if A[i].id == idP {
			return true
		}
	}
	return false
}

func pilihUbahHapusData() {
	var pilih int

	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("┃         UBAH / HAPUS DATA PEMINJAM        ┃")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Println("┃ [1] Ubah Data Peminjam                    ┃")
	fmt.Println("┃ [2] Hapus Data Peminjam                   ┃")
	fmt.Println("┃ [0] Kembali ke Menu Utama                 ┃")
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("Pilih No ➝  ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		ubahData(&data, nData)
	case 2:
		hapusData(&data, &nData)
	}
}

func ubahData(A *tabPinjaman, n int) {
	var i int
	var id string
	var found bool = false

	cetakData(data, nData)
	fmt.Print("Masukkan ID yang Ingin Diubah: ")
	fmt.Scan(&id)

	for i = 0; i < n && found == false; i++ {
		if A[i].id == id {
			fmt.Println("Masukkan Data Baru Untuk Mengganti.")

			fmt.Print("Nama Peminjam          : ")
			fmt.Scan(&A[i].nama)

			fmt.Print("Jumlah Pinjaman (Rp)   : ")
			fmt.Scan(&A[i].pinjaman)

			fmt.Print("Tenor Pinjaman (Bulan) : ")
			fmt.Scan(&A[i].tenor)

			found = true
		}
	}
	if found == false {
		fmt.Println("ID Tidak Ditemukan!")
	} else {
		fmt.Println("Data Berhasil Diubah!")
	}
}

func hapusData(A *tabPinjaman, n *int) {
	var i, j int
	var id string
	var found bool = false

	cetakData(data, nData)
	fmt.Print("Masukkan ID yang Ingin Dihapus: ")
	fmt.Scan(&id)

	for i = 0; i < *n && found == false; i++ {
		if A[i].id == id {
			for j = i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			found = true
			*n = *n - 1
		}
	}
	if found == false {
		fmt.Println("ID Tidak Ditemukan")
	} else {
		fmt.Println("Data Berhasil Dihapus")
	}
}

func pilihSort() {
	var pilih int

	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("┃          PENGURUTAN DATA PEMINJAM         ┃")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Println("┃ [1] Data Terurut Menaik - Pinjaman        ┃")
	fmt.Println("┃ [2] Data Terurut Menaik - Tenor           ┃")
	fmt.Println("┃ [3] Data Terurut Menurun - Pinjaman       ┃")
	fmt.Println("┃ [4] Data Terurut Menurun - Tenor          ┃")
	fmt.Println("┃ [0] Kembali ke Menu Utama                 ┃")
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("Pilih No ➝  ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		insertionSortPinjaman(&data, nData)
	case 2:
		insertionSortTenor(&data, nData)
	case 3:
		selectionSortPinjaman(&data, nData)
	case 4:
		selectionSortTenor(&data, nData)
	}
}

func insertionSortPinjaman(A *tabPinjaman, n int) {
	var pass, i int
	var temp pinjaman
	pass = 1
	for pass < n {
		i = pass
		temp = A[pass]
		for i > 0 && temp.pinjaman < A[i-1].pinjaman {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
	cetakData(data, nData)
	fmt.Println("Data Peminjam Telah Diurutkan Menaik - Pinjaman.")
}

func insertionSortTenor(A *tabPinjaman, n int) {
	var pass, i int
	var temp pinjaman
	pass = 1
	for pass < n {
		i = pass
		temp = A[pass]
		for i > 0 && temp.tenor < A[i-1].tenor {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
	cetakData(data, nData)
	fmt.Println("Data Peminjam Telah Diurutkan Menaik - Tenor.")
}

func selectionSortPinjaman(A *tabPinjaman, n int) {
	var pass, idx, i int
	var temp pinjaman
	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].pinjaman > A[idx].pinjaman {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
	cetakData(data, nData)
	fmt.Println("Data Peminjam Telah Diurutkan Menurut - Pinjaman.")
}

func selectionSortTenor(A *tabPinjaman, n int) {
	var pass, idx, i int
	var temp pinjaman
	pass = 1
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if A[i].tenor > A[idx].tenor {
				idx = i
			}
			i++
		}
		temp = A[pass-1]
		A[pass-1] = A[idx]
		A[idx] = temp
		pass++
	}
	cetakData(data, nData)
	fmt.Println("Data Peminjam Telah Diurutkan Menurun - Tenor.")
}

//tinggal edit func kebawah + status pembayaran belum ada :)
func hitungBunga(A *tabPinjaman, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Printf("Masukkan Suku Bunga Untuk Data ke-%d (%%): ", i+1)
		fmt.Scan(&A[i].bunga)
	}

	for i = 0; i < n; i++ {
		A[i].bunga = (A[i].bunga + 100) / 100
		A[i].tBunga = A[i].bunga * float64(A[i].pinjaman)
		A[i].kredit = A[i].tBunga / float64(A[i].tenor)
	}
	for i = 0; i < n; i++ {
		fmt.Printf("%.0f %.0f \n", A[i].tBunga, A[i].kredit)
	}
}

//func pilihCari udh diedit
func pilihCari() {
	var pilih int

	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("┃           PENCARIAN DATA PEMINJAM         ┃")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Println("┃ [1] Cari ID Peminjam                      ┃")
	fmt.Println("┃ [2] Cari                                  ┃")
	fmt.Println("┃ [3] Cari Pinjaman Terendah                ┃")
	fmt.Println("┃ [4] Cari Pinjaman Tertinggi               ┃")
	fmt.Println("┃ [0] Back                                  ┃")
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
	fmt.Print("Pilih No ➝  ")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		cariSequentialData(&data, nData)
	case 2:
		cariBinaryData(&data, nData)
	case 3:
		nilaiMin(data, nData)
	case 4:
		nilaiMax(data, nData)
	}
}

func cariSequentialData(A *tabPinjaman, n int) {
	var i int
	var id string
	var found bool = false

	cetakData(data, nData)
	fmt.Print("Masukkan ID yang ingin dicari: ")
	fmt.Scan(&id)

	for i = 0; i < n && found == false; i++ {
		if A[i].id == id {
			found = true
			fmt.Println("ID:", A[i].id)
			fmt.Println("Nama:", A[i].nama)
			fmt.Println("Pinjaman:", A[i].pinjaman)
			fmt.Println("Tenor:", A[i].tenor)
			fmt.Printf("Total pembayaran: %.2f\n", A[i].tBunga)
			fmt.Printf("Cicilan per bulan: %.2f\n", A[i].kredit)
		}
	}
	if found == false {
		fmt.Println("ID tidak ditemukan")
	}
}

func cariBinaryData(A *tabPinjaman, n int) {
	var pinjaman int
	var kiri, kanan, tengah int
	var found bool = false

	insertionSortPinjaman(A, n)

	fmt.Print("Masukkan jumlah pinjaman yang ingin dicari: ")
	fmt.Scan(&pinjaman)

	kiri = 0
	kanan = n - 1

	for kiri <= kanan && found == false {
		tengah = (kiri + kanan) / 2
		if A[tengah].pinjaman == pinjaman {
			fmt.Println("ID:", A[tengah].id)
			fmt.Println("Nama:", A[tengah].nama)
			fmt.Println("Pinjaman:", A[tengah].pinjaman)
			fmt.Println("Tenor:", A[tengah].tenor)
			fmt.Printf("Total pembayaran: %.0f\n", A[tengah].tBunga)
			fmt.Printf("Cicilan per bulan: %.0f\n", A[tengah].kredit)
			found = true
		} else if pinjaman < A[tengah].pinjaman {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
}

func nilaiMax(A tabPinjaman, n int) {
	var i, idxMax int
	for i = 0; i < n; i++ {
		if A[i].tBunga > A[idxMax].tBunga {
			idxMax = i
		}
	}
	fmt.Println("Data dengan total pembayaran tertinggi:")
	fmt.Println("ID:", A[idxMax].id)
	fmt.Println("Nama:", A[idxMax].nama)
	fmt.Println("Pinjaman:", A[idxMax].pinjaman)
	fmt.Println("Tenor:", A[idxMax].tenor)
	fmt.Printf("Total pembayaran: %.0f\n", A[idxMax].tBunga)
	fmt.Printf("Cicilan per bulan: %.0f\n", A[idxMax].kredit)
}

func nilaiMin(A tabPinjaman, n int) {
	var i, idxMin int
	for i = 0; i < n; i++ {
		if A[i].tBunga < A[idxMin].tBunga {
			idxMin = i
		}
	}
	fmt.Println("Data dengan total pembayaran terendah:")
	fmt.Println("ID:", A[idxMin].id)
	fmt.Println("Nama:", A[idxMin].nama)
	fmt.Println("Pinjaman:", A[idxMin].pinjaman)
	fmt.Println("Tenor:", A[idxMin].tenor)
	fmt.Printf("Total pembayaran: %.0f\n", A[idxMin].tBunga)
	fmt.Printf("Cicilan per bulan: %.0f\n", A[idxMin].kredit)
}

func cetakData(A tabPinjaman, n int) {
	var i int

	fmt.Println("Data Saat Ini")
	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Printf("┃ %-5s ┃ %-25s ┃ %-13s ┃ %-5s ┃\n", "ID", "Nama", "Pinjaman", "Tenor")
	fmt.Println("┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃")
	for i = 0; i < n; i++ {
		fmt.Printf("┃ %-5s ┃ %-25s ┃ %-13d ┃ %-5d ┃\n", A[i].id, A[i].nama, A[i].pinjaman, A[i].tenor)
	}
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
}

func cetakKredit(A tabPinjaman, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Printf("%s %s %d %d %.0f %.0f \n", A[i].id, A[i].nama, A[i].pinjaman, A[i].tenor, A[i].tBunga, A[i].kredit)
	}
}

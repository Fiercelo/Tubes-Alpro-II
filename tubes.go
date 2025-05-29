package main

import "fmt"

const NMAX int = 10

type pinjaman struct {
	id         string
	nama       string
	pinjaman   int
	tenor      int
	bunga      float64
	tBunga     float64
	kredit     float64
	totalBayar float64
	status     string
	sisa       float64
}
type tabPinjaman [NMAX]pinjaman

var data tabPinjaman
var nData int

func main() {
	fmt.Println("⎛⎝ ≽ > ⩊ < ≼ ⎠⎞      ﷽     ⎛⎝ ≽ > ⩊ < ≼ ⎠⎞")
	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("┃   Selamat Datang di Aplikasi Pinjaman  ┃")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Println("┃   Aplikasi ini Mengsimulasikan Sistem  ┃")
	fmt.Println("┃     Pinjaman dan Kredit Sederhana      ┃")
	fmt.Println("┃       Tekan ENTER Untuk Memulai        ┃")
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
		fmt.Println("┃ [6] Status Pembayaran                  ┃")
		fmt.Println("┃ [7] Tampilkan Laporan                  ┃")
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
			cetakKredit(data, nData)
			bayarCicilan(&data, nData)
			statusPembayaran(&data, nData)
		case 7:
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
	var i, jumlah, max int
	var idP string

	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("┃         TAMBAH DATA PEMINJAM BARU         ┃")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Println("┃ PENJELASAN                                ┃")
	fmt.Println("┃ Setiap peminjam punya:                    ┃")
	fmt.Println("┃ ID unik, nama, jumlah pinjaman, dan tenor ┃")
	fmt.Println("┃ NOTE: Maksimal 5 huruf/angka untuk ID dan ┃")
	fmt.Println("┃       maksimal 2 kata untuk nama          ┃")
	fmt.Println("┃                                           ┃")
	fmt.Println("┃ Contoh                                    ┃")
	fmt.Println("┃  ID Unik                : CL01            ┃")
	fmt.Println("┃  Nama Peminjam          : Dedi_Gusnaldi   ┃")
	fmt.Println("┃  Jumlah Pinjaman (Rp)   : 10000000        ┃")
	fmt.Println("┃  Tenor Pinjaman (bulan) : 12              ┃")
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")

	fmt.Print("Jumlah Data yang Ingin Dimasukkan (Max 10): ")
	fmt.Scan(&jumlah)
	max = NMAX - *n
	if max > 0 {
		if jumlah > max {
			jumlah = max
			fmt.Println("Kapasitas Maksimal Tercapai. Data yang Ditambahkan Dibatasi Menjadi ", max)
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
					fmt.Println("ID Sudah Digunakan, Silakan Masukkan ID Lain!")
				}
			}

			fmt.Print("Nama peminjam          : ")
			fmt.Scan(&A[*n].nama)

			fmt.Print("Jumlah Pinjaman (Rp)   : ")
			fmt.Scan(&A[*n].pinjaman)

			fmt.Print("Tenor Pinjaman (Bulan) : ")
			fmt.Scan(&A[*n].tenor)

			*n++
		}
		fmt.Println("\nData Berhasil Ditambahkan!")
	} else {
		fmt.Println("Data Tidak Dapat Ditambahkan, Kapasitas Sudah Penuh!")
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

	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("┃         MENGHITUNG BUNGA & CICILAN        ┃")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Println("┃ Masukkan Suku Bunga Untuk Tiap Peminjam.  ┃")
	fmt.Println("┃ Bunga Akan Dihitung Terhadap Jumlah       ┃")
	fmt.Println("┃ Pinjaman dan Tenor Dalam Bulan.           ┃")
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")

	for i = 0; i < n; i++ {
		fmt.Printf(" Masukkan Suku Bunga Untuk Data ke-%d (%%): ", i+1)
		fmt.Scan(&A[i].bunga)
	}

	for i = 0; i < n; i++ {
		A[i].bunga = (A[i].bunga + 100) / 100
		A[i].tBunga = A[i].bunga * float64(A[i].pinjaman)
		A[i].kredit = A[i].tBunga / float64(A[i].tenor)
	}
	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Printf("┃ %-5s ┃ %-22s ┃\n", "Total Pembayaran", "Cicilan per Bulan")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	for i = 0; i < n; i++ {
		fmt.Printf("┃ %-16.0f ┃ %-22.0f ┃\n", A[i].tBunga, A[i].kredit)
	}
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
}

//func pilihCari udh diedit
func pilihCari() {
	var pilih int

	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("┃           PENCARIAN DATA PEMINJAM         ┃")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Println("┃ [1] Cari Data Peminjam dengan ID          ┃")
	fmt.Println("┃ [2] Cari Data Peminjam dengan Pinjaman    ┃")
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
	fmt.Print("Masukkan ID: ")
	fmt.Scan(&id)

	for i = 0; i < n && found == false; i++ {
		if A[i].id == id {
			found = true
			fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
			fmt.Printf("┃ %-20s : %-18v ┃\n", "ID: ", A[i].id)
			fmt.Printf("┃ %-20s : %-18v ┃\n", "Nama: ", A[i].nama)
			fmt.Printf("┃ %-20s : %-18v ┃\n", "Pinjaman: ", A[i].pinjaman)
			fmt.Printf("┃ %-20s : %-18v ┃\n", "Tenor: ", A[i].tenor)
			fmt.Printf("┃ %-20s : %-18.0f ┃\n", "Total Pembayaran: ", A[i].tBunga)
			fmt.Printf("┃ %-20s : %-18.0f ┃\n", "Cicilan per Bulan: ", A[i].kredit)
			fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
		}
	}
	if found == false {
		fmt.Println("ID Tidak Ditemukan")
	}
}

func cariBinaryData(A *tabPinjaman, n int) {
	var pinjaman int
	var kiri, kanan, tengah int
	var found bool = false

	insertionSortPinjaman(A, n)

	fmt.Print("Masukkan Pinjaman: ")
	fmt.Scan(&pinjaman)

	kiri = 0
	kanan = n - 1

	for kiri <= kanan && found == false {
		tengah = (kiri + kanan) / 2
		if A[tengah].pinjaman == pinjaman {
			fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
			fmt.Printf("┃ %-20s : %-18v ┃\n", "ID: ", A[tengah].id)
			fmt.Printf("┃ %-20s : %-18v ┃\n", "ID: ", A[tengah].nama)
			fmt.Printf("┃ %-20s : %-18v ┃\n", "Nama: ", A[tengah].pinjaman)
			fmt.Printf("┃ %-20s : %-18v ┃\n", "Pinjaman: ", A[tengah].tenor)
			fmt.Printf("┃ %-20s : %-18v ┃\n", "Tenor: ", A[tengah].tBunga)
			fmt.Printf("┃ %-20s : %-18.0f ┃\n", "Total Pembayaran: ", A[tengah].tBunga)
			fmt.Printf("┃ %-20s : %-18.0f ┃\n", "Cicilan per Bulan: ", A[tengah].kredit)
			fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
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
	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Printf("┃ \033[31m%-42s\033[0m┃\n", "  Data Dengan Total Pembayaran Tertinggi")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Printf("┃ %-20s : %-18v ┃\n", "ID: ", A[idxMax].id)
	fmt.Printf("┃ %-20s : %-18v ┃\n", "Nama: ", A[idxMax].nama)
	fmt.Printf("┃ %-20s : %-18v ┃\n", "Pinjaman: ", A[idxMax].pinjaman)
	fmt.Printf("┃ %-20s : %-18v ┃\n", "Tenor: ", A[idxMax].tenor)
	fmt.Printf("┃ %-20s : %-18.0f ┃\n", "Total Pembayaran: ", A[idxMax].tBunga)
	fmt.Printf("┃ %-20s : %-18.0f ┃\n", "Cicilan per Bulan: ", A[idxMax].kredit)
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
}

func nilaiMin(A tabPinjaman, n int) {
	var i, idxMin int
	for i = 0; i < n; i++ {
		if A[i].tBunga < A[idxMin].tBunga {
			idxMin = i
		}
	}
	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Printf("┃ \033[34m%-42s\033[0m┃\n", " Data Dengan Total Pembayaran Terendah")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	fmt.Printf("┃ %-20s : %-18v ┃\n", "ID: ", A[idxMin].id)
	fmt.Printf("┃ %-20s : %-18v ┃\n", "Nama: ", A[idxMin].nama)
	fmt.Printf("┃ %-20s : %-18v ┃\n", "Pinjaman: ", A[idxMin].pinjaman)
	fmt.Printf("┃ %-20s : %-18v ┃\n", "Tenor: ", A[idxMin].tenor)
	fmt.Printf("┃ %-20s : %-18.0f ┃\n", "Total Pembayaran: ", A[idxMin].tBunga)
	fmt.Printf("┃ %-20s : %-18.0f ┃\n", "Cicilan per Bulan: ", A[idxMin].kredit)
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
}

func bayarCicilan(A *tabPinjaman, n int) {
	var id string
	var i int
	var ketemu bool = false

	fmt.Print("Masukkan ID Nasabah: ")
	fmt.Scan(&id)

	for i = 0; i < n; i++ {
		if A[i].id == id && !ketemu {
			ketemu = true
			fmt.Printf("Masukkan Jumlah Cicilan yang Dibayar Oleh %s: ", A[i].nama)
			fmt.Scan(&A[i].totalBayar)

			if A[i].totalBayar >= A[i].tBunga {
				A[i].status = "Lunas"
			} else {
				A[i].sisa = A[i].tBunga - A[i].totalBayar
				A[i].status = "Belum Lunas"
			}
			break
		}
	}

	if !ketemu {
		fmt.Println("ID Tidak Ditemukan")
	}
}

func statusPembayaran(A *tabPinjaman, n int) {
	var i int
	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Println("┃ ID   ┃ Nama         ┃ Pinjaman   ┃ Tenor ┃ Total Bayar  ┃ Cicilan    ┃ Sisa Bayar     ┃     Status     ┃")
	fmt.Println("┣━━━━━━┃━━━━━━━━━━━━━━┃━━━━━━━━━━━━┃━━━━━━━┃━━━━━━━━━━━━━━┃━━━━━━━━━━━━┃━━━━━━━━━━━━━━━━┫━━━━━━━━━━━━━━━━┫")

	for i = 0; i < n; i++ {
		fmt.Printf("┃ %-4s ┃ %-12s ┃ %-10d ┃ %-5d ┃ %-12.0f ┃ %-10.0f ┃ %-14.0f ┃ %-14s ┃\n",
			A[i].id, A[i].nama, A[i].pinjaman, A[i].tenor, A[i].tBunga, A[i].kredit, A[i].sisa, A[i].status)
	}

	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
}

func cetakData(A tabPinjaman, n int) {
	var i int

	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Printf("┃ %-5s ┃ %-25s ┃ %-13s ┃ %-5s ┃\n", "ID", "Nama", "Pinjaman", "Tenor")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	for i = 0; i < n; i++ {
		fmt.Printf("┃ %-5s ┃ %-25s ┃ %-13d ┃ %-5d ┃\n", A[i].id, A[i].nama, A[i].pinjaman, A[i].tenor)
	}
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")
}

func cetakKredit(A tabPinjaman, n int) {
	var i int

	fmt.Println("┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓")
	fmt.Printf("┃ %-5s ┃ %-25s ┃ %-13s ┃ %-5s ┃ %-16s ┃ %-17s ┃\n", "ID", "Nama", "Pinjaman", "Tenor", "Total Pembayaran", "Cicilan per Bulan")
	fmt.Println("┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫")
	for i = 0; i < n; i++ {
		fmt.Printf("┃ %-5s ┃ %-25s ┃ %-13d ┃ %-5d ┃ %-16.0f ┃ %-17.0f ┃\n", A[i].id, A[i].nama, A[i].pinjaman, A[i].tenor, A[i].tBunga, A[i].kredit)
	}
	fmt.Println("┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛")

}

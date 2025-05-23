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
	menu()
}

func menu() {
	var pilih int

	for {
		fmt.Println("┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃")
		fmt.Println("┃               PINJAMAN BANK              ┃")
		fmt.Println("┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃")
		fmt.Println("┃1. Menambahkan Data Peminjam              ┃")
		fmt.Println("┃2. Mengubah atau Menghapus Data Peminjam  ┃")
		fmt.Println("┃3. Mengurutkan Daftar Peminjam            ┃")
		fmt.Println("┃4. Menghitung Data                        ┃")
		fmt.Println("┃5. Cari Data                              ┃")
		fmt.Println("┃6. Tampilkan Laporan                      ┃")
		fmt.Println("┃0. EXIT                                   ┃")
		fmt.Print(" Pilih No ➝  ")

		fmt.Scan(&pilih)
		fmt.Println("┃━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┃")

		switch pilih {
		case 1:
			tambahData(&data, &nData)
		case 2:
			if nData == 0 {
				fmt.Println("Masukkan data peminjam terlebih dahulu")
			} else {
				pilihUbahHapusData()
			}
		case 3:
			if nData == 0 {
				fmt.Println("Masukkan data peminjam terlebih dahulu")
			} else {
				pilihSort()
			}
		case 4:
			if nData == 0 {
				fmt.Println("Masukkan data peminjam terlebih dahulu")
			} else {
				hitungBunga(&data, nData)
			}
		case 5:
			if nData == 0 {
				fmt.Println("Masukkan data peminjam terlebih dahulu")
			} else {
				pilihCari()
			}
		case 6:
			cetakData(data, nData)
		default:
			fmt.Println("Pilihan tidak tersedia. Coba lagi.")
		}
		if pilih == 0 {
			break
		}
	}
}

func tambahData(A *tabPinjaman, n *int) {
	var i, jumlah int
	var idP string

	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("           TAMBAH DATA PEMINJAM BARU           ")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Print("Jumlah data yang ingin dimasukkan (max 10): ")
	fmt.Scan(&jumlah)

	if jumlah > NMAX {
		jumlah = NMAX
		fmt.Println("Kapasitas maksimal tercapai. Data yang ditambahkan dibatasi menjadi 10.")
	}

	for i = 0; i < jumlah; i++ {
		fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
		fmt.Printf("Data peminjam ke-%d\n", *n+1)

		for {
			fmt.Print("Masukkan ID unik: ")
			fmt.Scan(&idP)
			if !IDSama(*A, *n, idP) {
				A[*n].id = idP
				break
			} else {
				fmt.Println("ID sudah digunakan, Silakan masukkan ID lain.")
			}
		}

		fmt.Print("Nama peminjam: ")
		fmt.Scan(&A[*n].nama)

		fmt.Print("Jumlah pinjaman (Rp): ")
		fmt.Scan(&A[*n].pinjaman)

		fmt.Print("Tenor pinjaman (bulan): ")
		fmt.Scan(&A[*n].tenor)

		*n++
	}
	fmt.Println("Data berhasil ditambahkan!")
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

	fmt.Println("Pilih ubah data atau hapus data             ")
	fmt.Println("1. Mengubah data peminjam                   ")
	fmt.Println("2. Menghapus data peminjam                  ")
	fmt.Println("0. BACK                                     ")
	fmt.Print("Pilih No ➝ ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		ubahData(&data, nData)
	} else if pilih == 2 {
		hapusData(&data, &nData)
	}
}

func ubahData(A *tabPinjaman, n int) {
	var i int
	var id string
	var found bool = false

	cetakData(data, nData)
	fmt.Print("Masukkan ID yang ingin diubah: ")
	fmt.Scan(&id)

	for i = 0; i < n && found == false; i++ {
		if A[i].id == id {
			fmt.Println("Ganti data yang ingin diubah")

			fmt.Print("Nama: ")
			fmt.Scan(&A[i].nama)

			fmt.Print("Pinjaman: ")
			fmt.Scan(&A[i].pinjaman)

			fmt.Print("Tenor (bulan): ")
			fmt.Scan(&A[i].tenor)

			found = true
		}
	}
	if found == false {
		fmt.Println("ID tidak ditemukan")
	} else {
		fmt.Println("Data berhasil diubah")
	}
}

func hapusData(A *tabPinjaman, n *int) {
	var i, j int
	var id string
	var found bool = false

	cetakData(data, nData)
	fmt.Print("Masukkan ID yang ingin dihapus: ")
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
		fmt.Println("ID tidak ditemukan")
	} else {
		fmt.Println("Data berhasil dihapus")
	}
}

func pilihCari() {
	var pilih int

	fmt.Println("Pilih pencarian                             ")
	fmt.Println("1. Cari ID peminjam                         ")
	fmt.Println("2. Cari pinjaman terendah                   ")
	fmt.Println("3. Cari pinjaman tertinggi                  ")
	fmt.Println("0. BACK                                     ")
	fmt.Print("Pilih No ➝  ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		cariData(&data, nData)
	} else if pilih == 2 {
		nilaiMin(data, nData)
	} else if pilih == 3 {
		nilaiMax(data, nData)
	}

}

func cariData(A *tabPinjaman, n int) {
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

func pilihSort() {
	var pilih int

	fmt.Println("Pilih metode pengurutan                     ")
	fmt.Println("1. Data terurut menaik berdasarkan pinjaman ")
	fmt.Println("2. Data terurut menaik berdasarkan tenor    ")
	fmt.Println("3. Data terurut menurun berdasarkan pinjaman")
	fmt.Println("4. Data terurut menurun berdasarkan tenor   ")
	fmt.Print("Pilih No ➝  ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		insertionSortPinjaman(&data, nData)
	} else if pilih == 2 {
		insertionSortTenor(&data, nData)
	} else if pilih == 3 {
		selectionSortPinjaman(&data, nData)
	} else if pilih == 4 {
		selectionSortTenor(&data, nData)
	}
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
}

func cetakData(A tabPinjaman, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Printf("%s %s %d %d \n", A[i].id, A[i].nama, A[i].pinjaman, A[i].tenor)
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

func hitungBunga(A *tabPinjaman, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Printf("Masukkan suku bunga untuk data ke-%d (%%): ", i+1)
		fmt.Scan(&A[i].bunga)
	}

	for i = 0; i < n; i++ {
		A[i].bunga = (A[i].bunga + 100) / 100
		A[i].tBunga = ((A[i].bunga + 100) / 100) * float64(A[i].pinjaman)
		A[i].kredit = A[i].tBunga / float64(A[i].tenor)
	}
	for i = 0; i < n; i++ {
		fmt.Printf("%s %s %d %d %.0f %.0f \n", A[i].id, A[i].nama, A[i].pinjaman, A[i].tenor, A[i].tBunga, A[i].kredit)
	}
}

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

func plihUbahHapusData() {
	var pilih int

	fmt.Println("Pilih ubah data atau hapus data")
	fmt.Println("1. Mengubah data peminjam")
	fmt.Println("2. Menghapus data peminjam")
	fmt.Println("0. BACK")
	fmt.Print("Pilih No ➝ ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		ubahData(&data, nData)
	} else if pilih == 2 {
		hapusData(&data, &nData)
	}
}

func tambahData(A *tabPinjaman, n *int) {
	var i, jumlah int

	fmt.Print("Jumlah data yang ingin dimasukkan: ")
	fmt.Scan(&jumlah)
	if jumlah > NMAX {
		jumlah = NMAX
	}
	for i = 0; i < jumlah; i++ {
		fmt.Println("====================================")
		fmt.Printf("Tambahkan data peminjam ke-%d\n", *n+1)

		fmt.Print("ID: ")
		fmt.Scan(&A[*n].id)

		fmt.Print("Nama peminjam: ")
		fmt.Scan(&A[*n].nama)

		fmt.Print("Jumlah pinjaman (Rp): ")
		fmt.Scan(&A[*n].pinjaman)

		fmt.Print("Tenor (bulan): ")
		fmt.Scan(&A[*n].tenor)

		*n = *n + 1
	}
	fmt.Println("Data berhasil ditambahkan")
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

func cariData(A *tabPinjaman, n int) {
	var i int
	var id string
	var found bool = false

	fmt.Print("Masukkan ID yang ingin dicari: ")
	fmt.Scan(&id)

	for i = 0; i < n && found == false; i++ {
		if A[i].id == id {
			found = true
			fmt.Println("ID:", A[i].id)
			fmt.Println("Nama:", A[i].nama)
			fmt.Println("Pinjaman:", A[i].pinjaman)
			fmt.Println("Tenor:", A[i].tenor)
		}
	}
	if found == false {
		fmt.Println("ID tidak ditemukan")
	}
}

func pilihSort() {
	var pilih int

	fmt.Println("Pilih metode pengurutan")
	fmt.Println("1. Data terurut menaik berdasarkan pinjaman")
	fmt.Println("2. Data terurut menaik berdasarkan tenor")
	fmt.Println("3. Data terurut menurun berdasarkan pinjaman")
	fmt.Println("4. Data terurut menurun berdasarkan tenor")
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
}

func cetakData(A tabPinjaman, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Printf("%s %s %d %d \n", A[i].id, A[i].nama, A[i].pinjaman, A[i].tenor)
	}
}

func nilaiMax(A tabPinjaman, n int) int {
	var i, idxMax int
	for i = 0; i < n; i++ {
		if A[i].pinjaman > A[idxMax].pinjaman {
			idxMax = i
		}
	}
	return idxMax
}

func nilaiMin(A tabPinjaman, n int) int {
	var i, idxMin int
	for i = 0; i < n; i++ {
		if A[i].pinjaman < A[idxMin].pinjaman {
			idxMin = i
		}
	}
	return idxMin
}

func hitungBunga(A *tabPinjaman, n int) {
	var i int
	for i = 0; i < n; i++ {
		A[i].bunga = (A[i].bunga + 100) / 100
		A[i].tBunga = A[i].bunga * float64(A[i].pinjaman)
		A[i].kredit = A[i].tBunga / float64(A[i].tenor)
	}
}

func menu() {
	var pilih int

	for {
		fmt.Println("========================================")
		fmt.Println("||           PINJAMAN BANK            ||")
		fmt.Println("========================================")
		fmt.Println("1. Menambahkan Data Peminjam            ")
		fmt.Println("2. Mengubah atau Menghapus Data Peminjam")
		fmt.Println("3. Menghitung Data                      ")
		fmt.Println("4. Cari Data                            ")
		fmt.Println("5. Mengurutkan Daftar Peminjam          ")
		fmt.Println("8. Tampilkan Laporan                    ")
		fmt.Println("0. EXIT                                 ")
		fmt.Print("Pilih No ➝  ")

		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahData(&data, &nData)
		case 2:
			if nData == 0 {
				fmt.Println("Masukkan data peminjam terlebih dahulu")
			} else {
				plihUbahHapusData()
			}
		case 3:
			hitungBunga(&data, nData)
		case 4:
			cariData(&data, nData)
		case 5:
			pilihSort()
		case 8:
			cetakData(data, nData)
		}
		if pilih == 0 {
			break
		}
	}
}

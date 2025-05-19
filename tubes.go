package main

import "fmt"

const NMAX int = 100

type pinjaman struct {
	id         int
	nama       string
	pinjaman   int
	tenor      int
	bunga      int
	jenisBunga string
}
type tabPinjaman [NMAX]pinjaman

func main() {
	menu()
}

func tambahData(A *tabPinjaman, n *int) {
	var i, jumlah int

	fmt.Print("Jumlah data yang ingin dimasukkan: ")
	fmt.Scan(&jumlah)
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

		fmt.Print("Suku bunga (%): ")
		fmt.Scan(&A[*n].bunga)

		fmt.Print("Jenis bunga (Tetap/Variabel): ")
		fmt.Scan(&A[*n].jenisBunga)

		*n = *n + 1
	}
	fmt.Println("Data berhasil ditambahkan")
}

func ubahData(A *tabPinjaman, n int) {
	var i int
	var id int
	var found bool = false

	fmt.Print("Masukkan id yang ingin diubah: ")
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

			fmt.Print("Suku bunga (%): ")
			fmt.Scan(&A[i].bunga)

			fmt.Print("Jenis bunga (Tetap/Variabel): ")
			fmt.Scan(&A[i].jenisBunga)

			found = true
		}
	}
	if found == false {
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Println("Data berhasil diubah")
	}
}

func hapusData(A *tabPinjaman, n *int) {
	var i, j, id int
	var found bool = false

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
		fmt.Println("Data tidak ditemukan")
	} else {
		fmt.Println("Data berhasil dihapus")
	}
}

func selectionSort(A *tabPinjaman, n int) {
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

func insertionSort(A *tabPinjaman, n int) {
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

func cetakData(A tabPinjaman, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Printf("%s %d %d \n", A[i].nama, A[i].pinjaman, A[i].tenor)
	}
}

func menu() {
	var data tabPinjaman
	var nData, pilih int

	for {
		fmt.Println("========================================")
		fmt.Println("||          PINJAMAN ONLINE           ||")
		fmt.Println("========================================")
		fmt.Println("1. Tambah Data                          ")
		fmt.Println("2. Ubah Data                            ")
		fmt.Println("3. Hapus Data                           ")
		fmt.Println("4. Hitung Data                          ")
		fmt.Println("5. Cari Data                            ")
		fmt.Println("6. Mengurutkan Data Secara Menaik       ")
		fmt.Println("7. Mengurutkan Data Secara Menurun      ")
		fmt.Println("8. Tampilkan Laporan                    ")
		fmt.Println("0. EXIT                                 ")
		fmt.Print("Pilih No ➝  ")

		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			tambahData(&data, &nData)
		case 2:
			ubahData(&data, nData)
		case 3:
			hapusData(&data, &nData)
		case 6:
			insertionSort(&data, nData)
		case 7:
			selectionSort(&data, nData)
		case 8:
			cetakData(data, nData)
		}
		if pilih == 0 {
			break
		}
	}
}

package main

import "fmt"

const NMAX int = 100

type pinjaman struct {
	nama     string
	pinjaman int
	tenor    int
}
type tabPinjaman [NMAX]pinjaman

func main() {
	menu()
}

func bacaData(A *tabPinjaman, n *int) {
	var i int

	fmt.Print("Jumlah data yang ingin dimasukkan: ")
	fmt.Scan(n)
	for i = 0; i < *n; i++ {
		fmt.Println("=========================")
		fmt.Printf("Data peminjam ke-%d\n", i+1)
		fmt.Print("Nama peminjam: ")
		fmt.Scan(&A[i].nama)

		fmt.Print("Jumlah pinjaman: ")
		fmt.Scan(&A[i].pinjaman)

		fmt.Print("Tenor (bulan): ")
		fmt.Scan(&A[i].tenor)
	}
}

func tambahData(A *tabPinjaman, n *int) {
	fmt.Println("Tambahkan data yang ingin dimasukkan")

	fmt.Print("Nama: ")
	fmt.Scan(&A[*n].nama)

	fmt.Print("Pinjaman: ")
	fmt.Scan(&A[*n].pinjaman)

	fmt.Print("Tenor(bulan): ")
	fmt.Scan(&A[*n].tenor)
	*n = *n + 1

	fmt.Println("Data berhasil ditambahkan")
}

func ubahData(A *tabPinjaman, n int) {
	var i int
	var nama string
	var found bool = false

	fmt.Print("Masukkan nama yang ingin diubah: ")
	fmt.Scan(&nama)

	for i = 0; i < n && found == false; i++ {
		if A[i].nama == nama {
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
		fmt.Println("Data tidak ditemukan")
	}

	fmt.Println("Data berhasil diubah")
}

func hapusData(A *tabPinjaman, n *int) {
	var i, j int
	var nama string
	var found bool = false

	fmt.Print("Masukkan nama yang ingin dihapus: ")
	fmt.Scan(&nama)

	for i = 0; i < *n && found == false; i++ {
		if A[i].nama == nama {
			for j = i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			fmt.Println("Data berhasil dihapus")
			found = true
			*n = *n - 1
		}
	}
	if found == false {
		fmt.Println("Data tidak ditemukan")
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
		fmt.Println("1. Masukkan Data                        ")
		fmt.Println("2. Tambah Data                          ")
		fmt.Println("3. Ubah Data                            ")
		fmt.Println("4. Hapus Data                           ")
		fmt.Println("5. Hitung Data                          ")
		fmt.Println("6. Cari Data                            ")
		fmt.Println("7. Mengurutkan Data                     ")
		fmt.Println("8. Tampilkan Laporan                    ")
		fmt.Println("0. EXIT                                 ")
		fmt.Print("Pilih No -> ")

		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			bacaData(&data, &nData)
		case 2:
			tambahData(&data, &nData)
		case 3:
			ubahData(&data, nData)
		case 4:
			hapusData(&data, &nData)
		case 7:
			cetakData(data, nData)
		}
		if pilih == 0 {
			break
		}
	}
}

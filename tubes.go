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

	fmt.Scan(n)
	for i = 0; i < *n; i++ {
		fmt.Scan(&A[i].nama, &A[i].pinjaman, &A[i].tenor)
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
		fmt.Println("5. Cari Data                            ")
		fmt.Println("6. Mengurutkan Data                     ")
		fmt.Println("7. Tampilkan Laporan                    ")
		fmt.Println("0. EXIT                                 ")
		fmt.Print("Pilih No -> ")

		fmt.Scan(&pilih)

		switch pilih {
		case 1:
			bacaData(&data, &nData)
		case 7:
			cetakData(data, nData)
		}
		if pilih == 0 {
			break
		}
	}
}

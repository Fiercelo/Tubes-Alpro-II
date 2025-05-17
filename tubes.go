package main
import "fmt"

const NMAX int = 100
type pinjaman struct {
  nama string
  pinjaman, tenor, pilih int
}
type tabPinjaman[NMAX] pinjaman

func main() {
	var data tabPinjaman
	var nData int

	menu() 
	fmt.Scan(&pilih)
	for pilih != 0 {

	}
}

func bacaData(A *tabPinjaman, n *int) {
	var i int
	
	fmt.Scan(n)
	for i = 0; i < *n; i++ {
		fmt.Scan(&A[i].nama, &A[i].pinjaman, &A[i].tenor)
	}

func cetakData(A tabPinjaman, n int) {
	var i int
	
	for i = 0; i < n; i++ {
		fmt.Printf("%s %d %d \n", A[i].nama, A[i].pinjaman, A[i].tenor)
	}
}

func menu() {
	fmt.Println("________________________________________")
	fmt.Println("|           PINJAMAN ONLINE            |")
	fmt.Println("________________________________________")
	fmt.Println("1. Masukkan Data Diri ")
	fmt.Println("2. Ubah Data Diri")
	fmt.Println("3. Tambah Data Diri")
	fmt.Println("4. Hapus Data Diri")
	fmt.Println("5.                ")
	fmt.Println("0. EXIT ")
	fmt.Println("Pilih No -> ")
}

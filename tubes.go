package main
import "fmt"

const NMAX int = 100
type pinjaman struct {
  nama string
  pinjaman, tenor int
}
type tabPinjaman[NMAX] pinjaman

func main() {
	var data tabPinjaman
	var nData int

	bacaData(&data, nData)
	cetakData(data, nData)
}

func bacaData() {
	var i int
	
	fmt.Scan(n)
	for i = 0; i < *n; i++ {
		fmt.Scan(&A[i].nama, &A[i].pinjaman, &A[i].tenor)
	}

func cetakData() {
	var i int
	
	for i = 0; i < n; i++ {
		fmt.Printf("%s %d %d \n", A[i].nama, A[i].pinjaman, A[i].tenor)
	}
}

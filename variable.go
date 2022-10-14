package main

import "fmt"

func main() {

	// set variable satu satu dengan deklarasi tipe data variable tersebut
	var nama1 string
	var umur1 int
	var alamat1 string

	nama1 = "Budi 1"
	umur1 = 31
	alamat1 = "Jakarta Barat"

	fmt.Println("Nama 1", nama1)
	fmt.Println("Umur 1", umur1)
	fmt.Println("Alamat 1", alamat1)
	fmt.Println("======================================")

	// membuat variable dengan deklarasi tipe data variable dan assign value
	var nama2 string = "Budi 2"
	var umur2 int = 32
	var alamat2 string = "Jakarta Selatan"

	fmt.Println("Nama 2", nama2)
	fmt.Println("Umur 2", umur2)
	fmt.Println("Alamat 2", alamat2)
	fmt.Println("======================================")

	var (
		nama   = "Budi"
		umur   = 25
		alamat = "Bojong Gede"
	)

	fmt.Println("Nama", nama)
	fmt.Println("Umur", umur)
	fmt.Println("Alamat", alamat)
}

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

	// membuat varible sekaligus dengan tipe data
	var (
		no_ktp     int
		no_telepon int
		bank       string
	)

	no_ktp = 12345677
	no_telepon = 87775365856
	bank = "Bank Negara Indonesia"

	fmt.Println("No KTP", no_ktp)
	fmt.Println("No Telepon", no_telepon)
	fmt.Println("Bank", bank)
	fmt.Println("======================================")

	/**
	 * Pembuatan group variable dengan langsung mengisi nilai variable
	 */
	var (
		nama   = "Budi"
		umur   = 25
		alamat = "Bojong Gede"
	)

	fmt.Println("Nama", nama)
	fmt.Println("Umur", umur)
	fmt.Println("Alamat", alamat)

	// membuat inline varible
	var (
		jenis_kelamin, tanggal_mulai, tanggal_akhir = "Laki-laki", "2022-01-01", "2022-10-01"
	)

	fmt.Println("Jenis Kelamin", jenis_kelamin)
	fmt.Println("Tanggal Mulai", tanggal_mulai)
	fmt.Println("Tanggal Akhir", tanggal_akhir)
}

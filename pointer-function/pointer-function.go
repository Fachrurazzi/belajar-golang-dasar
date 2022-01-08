package main

import "fmt"

type Mahasiswa struct {
	Nama, Npm, Alamat string
}

func changeAddressMahasiswa(mahasiswa *Mahasiswa) {
	mahasiswa.Alamat = "Banjarmasin"
}

func main() {
	mahasiswa := Mahasiswa{
		Nama:   "Razzi",
		Npm:    "16630942",
		Alamat: "Banjarbaru",
	}
	change := &mahasiswa
	changeAddressMahasiswa(change)
	fmt.Println(mahasiswa)
}

package main

import (
	"fmt"

	gosayhello "github.com/asepdwisaputra/go-say-hello/v5"
)

func main() {
	mahasiswa := gosayhello.Mahasiswa{
		Nama:   "Asep",
		Alamat: "Indo",
		NoHP:   0,
	}
	fmt.Println(mahasiswa)
	mahasiswa.SayHello("Irma")
}

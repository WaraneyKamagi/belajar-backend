package main

import "fmt"

type Alamat struct {
	Kota    string
	Negara  string
	KodePos int
}

type Person struct {
	Nama   string
	Alamat Alamat
}

func main() {
	person := Person{
		Nama: "Waraney",
		Alamat: Alamat{
			Kota:    "Manado",
			Negara:  "Indonesia",
			KodePos: 95111,
		},
	}
	fmt.Println(person.Nama,
		person.Alamat.Kota,
		person.Alamat.Negara,
		person.Alamat.KodePos)
}

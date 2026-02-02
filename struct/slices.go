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
	people := []Person{
		{
			Nama: "Waraney",
			Alamat: Alamat{
				Kota:    "Manado",
				Negara:  "Indonesia",
				KodePos: 95111,
			},
		},
		{
			Nama: "Budi",
			Alamat: Alamat{
				Kota:    "Jakarta",
				Negara:  "Indonesia",
				KodePos: 10110,
			},
		},
	}

	for i, ppl := range people {
		fmt.Println(i, ppl.Nama,
			ppl.Alamat.Kota,
			ppl.Alamat.Negara,
			ppl.Alamat.KodePos)
	}
}

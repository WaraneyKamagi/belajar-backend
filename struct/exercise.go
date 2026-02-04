package main

import "fmt"

type Laptop struct {
	Merk  string
	Specs Specs
}

type Specs struct {
	GPU         string
	Prosessor   string
	ClockSpeed  string
	RAM         int
	Penyimpanan string
}

func main() {
	laptop := []Laptop{
		{
			Merk: "Lenovo LOQ 13",
			Specs: Specs{
				GPU:         "NVIDIA GeForce RTX 4060",
				Prosessor:   "Intel Core i7-13700H",
				ClockSpeed:  "2.5 GHz",
				RAM:         16,
				Penyimpanan: "SSD 512GB",
			},
		},
		{
			Merk: "ASUS ROG Strix G16",
			Specs: Specs{
				GPU:         "NVIDIA GeForce RTX 4060",
				Prosessor:   "Intel Core i5-13450HX (10 inti) atau i7-13650HX",
				ClockSpeed:  "2.5 GHz",
				RAM:         16,
				Penyimpanan: "SSD 512GB",
			},
		},
	}
	for i, lpt := range laptop {
		fmt.Println(i, "Merk:", lpt.Merk,
			"GPU:", lpt.Specs.GPU,
			"Prosessor:", lpt.Specs.Prosessor,
			"ClockSpeed:", lpt.Specs.ClockSpeed,
			"RAM:", lpt.Specs.RAM,
			"Penyimpanan:", lpt.Specs.Penyimpanan)
	}
}

package main

import "fmt"

type MataKuliah struct {
	namaMatkul string
	namaDosen  string
	ruangKelas int
}

// Value receiver - read only: menampilkan informasi mata kuliah
func (m MataKuliah) Tampilkan() {
	fmt.Printf("Mata Kuliah: %s\nDosen: %s\nRuang: %d\n", m.namaMatkul, m.namaDosen, m.ruangKelas)
}

// Pointer receiver - mengubah informasi mata kuliah
func (m *MataKuliah) Ubah(namaMatkul, namaDosen string, ruangKelas int) {
	m.namaMatkul = namaMatkul
	m.namaDosen = namaDosen
	m.ruangKelas = ruangKelas
}

func main() {
	mat := MataKuliah{"Mobile development", "Sir Arga", 201}
	mat.Tampilkan()                                   // value receiver on value
	mat.Ubah("Sistem Pemrogramman", "Sir Rolly", 501) // pointer receiver on value (Go takes &mat)
	mat.Tampilkan()

	p := &MataKuliah{"Back-End Development", "Sir Enrico", 301}
	p.Tampilkan() // value receiver on pointer
	p.Ubah("Comnet", "Sir Wilson", 102)
	p.Tampilkan()

	q := &MataKuliah{"Aljabar", "Sir Sahulata", 303}
	q.Tampilkan() // value receiver on pointer
	q.Ubah("HCI", "Sir George", 101)
}

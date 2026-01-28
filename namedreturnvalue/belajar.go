package main

import "fmt"

func rectProps(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	// jika parameter tidak digunakan, akan mengembalikan nilai 0
	return // naked return - area dan perimeter otomatis dikembalikan
}
func main() {
	luas, keliling := rectProps(19.8, 5.6)
	fmt.Println("Area:", luas)
	fmt.Println("Perimeter:", keliling)
}

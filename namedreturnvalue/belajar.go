package main

import "fmt"

func calc(length, width float64) (hasil string, area float64) {
	area = length * width

	if area > 50 {
		hasil = "Besar"
	} else {
		hasil = "kecil"
	}
	return
}
func main() {

	size, luas := calc(5, 10.6)
	fmt.Println("Area:", size)
	fmt.Println("Luas:", luas)
}

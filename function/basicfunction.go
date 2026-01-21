package main

import "fmt"

func add(a int, b int) int {
	if result := a + b; result > 90 {
		fmt.Println("Hasilnya tinggi")
	}
	return a + b
}

func main() {
	result := add(50, 100)
	fmt.Println("Hasil penjumlahan:", result)
}

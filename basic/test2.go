package main

import "fmt"

func main() {
	if absen, nilai := 8, 91; absen > 7 {
		fmt.Println("WF")
	} else if nilai > 90 {
		fmt.Println("A")
	} else if nilai > 77 {
		fmt.Println("B")
	} else if nilai > 66 {
		fmt.Println("C")
	} else {
		fmt.Println("F")
	}
}

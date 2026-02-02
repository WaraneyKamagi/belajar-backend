package main

import "fmt"

type Student struct {
	Nama  string
	Umur  int
	Email string
}

func main() {
	student := Student{"Waraney", 20, "waraney@gmail.com"}
	student.Email = "waraney@universitasklabat.com"
	fmt.Println(student.Nama, student.Umur, student.Email)
}

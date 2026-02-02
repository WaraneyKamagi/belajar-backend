package main

import (
	"fmt"
	"multiplefiles/functions"
)

func triple(x int) int {
	return x * 3
}

func main() {
	result := functions.Apply(5, triple)
	fmt.Println(result)
}

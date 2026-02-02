package main

import (
	"fmt"
)

func apply(x int, f func(int) int) int {
	return f(x)
}

func triple(x int) int {
	return x * 3
}

func main() {
	result := apply(5, triple)
	fmt.Println(result)
}

package main

import "fmt"

func increment(x *int) {
	*x = *x + 10
}

func main() {
	num := 10
	increment(&num)
	fmt.Println(num)
}

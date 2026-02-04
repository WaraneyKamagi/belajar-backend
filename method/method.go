package main

import "fmt"

type Rectangle struct {
	area, length, width float32
}

// Value reciever - read only
func (r Rectangle) Area() {
	r.area = r.length * r.width
}

func (r *Rectangle) ScaleBy(factor float32) {
	r.length *= factor
	r.width *= factor
}

func main() {
	rect := Rectangle{0, 5, 10}
	rect.Area()     // Works: Value reciever on value
	rect.ScaleBy(2) // Works: pointer reciever on value (go takes &ract)
	fmt.Println(rect.area, rect.width, rect.length)
	p := &Rectangle{0, 5, 10}
	p.Area() // Works: Value reciever on pointer (Go dereferences)
	p.ScaleBy(2)
	fmt.Println(p.area, p.width, p.length)
}

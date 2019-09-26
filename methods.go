package main

import "fmt"

type rect struct {
	width, height int
	name          string
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) talk() string {
	return r.name
}

func (r rect) perim() int {
	return 2*r.height + 2*r.width
}

func main() {

	r := rect{10, 5, "Little a"}

	fmt.Println("area: ", r.area())

	fmt.Println("My name is: ", r.talk())

	rp := &r

	fmt.Println(rp.perim())

}

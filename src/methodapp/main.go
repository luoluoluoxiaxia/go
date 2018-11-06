package main

type Rectangle struct {
	width, height int
}

func (r Rectangle) area() int {
	return r.width * r.height
}

func main() {
	r1 := Rectangle{10,20}
	println(r1.area())
}

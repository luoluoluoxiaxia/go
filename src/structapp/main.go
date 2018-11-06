package main

type person struct {
	name string
	age  int
}

type student struct {
	id int
	person
}

func main() {
	p := person{"张三",10}

	p.age = 20

	println(p.name,p.age)

	s := student{}

	s.id = 10
	s.age = 20
	s.name = "哈哈"

	println(s.id,s.age,s.name)
}

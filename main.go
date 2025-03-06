package main

import "fmt"

type inter interface {
	Metod()
}
type Human struct {
	name string
	age  int
}

func (h Human) Metod() {
	fmt.Println("я человек", h.name, h.age)
}
func (h Human) Run() {
	fmt.Println("я побежал")
}

type Dag struct {
	name string
	age  int
}

func (d Dag) Metod() {
	fmt.Println("я утка", d.name, d.age)
}
func (d Dag) Run1() {
	fmt.Println("я плыву!")
}
func Res(i inter) {
	i.Metod()
	switch v := i.(type) {
	case Human:
		v.Run()
	case Dag:
		v.Run1()
	}
}
func main() {
	asd := Human{"max", 15}
	zxc := Dag{"gaga", 10}
	Res(asd)
	Res(zxc)
}

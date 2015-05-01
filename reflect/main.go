package main

import (
	"fmt"
	"reflect"
)

var _ = reflect.TypeOf

type Greeter interface {
	Greet(str string)
}

type (
	Porter struct {
		hotel string
	}
	Governor struct {
		city string
	}
)

func (p Porter) Greet(visitor string) {
	fmt.Println(visitor + ", welcome to The " + p.hotel + " hotel!")
}

func (g Governor) Greet(visitor string) {
	fmt.Println(visitor + ", welcome to " + g.city + " city!")
}

func main() {
	g0 := Greeter(Porter{"Vermington"})
	g1 := Greeter(Governor{"Makao"})

	g0.Greet("Sally")
	g1.Greet("Peter")

	reflect.ValueOf(g0).Interface().(Greeter).Greet("Sally")
	reflect.ValueOf(&g1).Elem().Interface().(Greeter).Greet("Peter")

	fmt.Println(reflect.ValueOf(g0).CanSet())
	fmt.Println(reflect.ValueOf(&g1).Elem().CanSet())
}

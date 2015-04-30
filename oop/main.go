package main

import (
	"fmt"
	"strconv"
)

type T struct {
	val int
}
func (t *T) Val() string {
	return strconv.Itoa(t.val)
}

type S struct {
	T
}

type V interface {
	Val() string
}

func main() {
	x := S{T{10}}
	var v V = x
	// s.val = 10
	fmt.Println("val = " + x.Val())
}

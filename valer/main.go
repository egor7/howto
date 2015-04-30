package main

import (
	"fmt"
	"strconv"
)

type (
	T struct {
		val int
	}
	S struct {
		T
	}
)

func (t *T) Val() string {
	return strconv.Itoa(t.val)
}

type Valer interface {
	Val() string
}

func main() {
	var s S
	s.val = 10
	var i Valer = &s
	fmt.Println("s.val = " + i.Val())
}

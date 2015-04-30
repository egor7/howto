package main

import (
	"bytes"
	"fmt"
	"io"
	"unsafe"
)

type iface struct {
	Type, Data unsafe.Pointer
}

func main() {
	var b = &bytes.Buffer{}
	var r io.Reader = b
	var i interface{} = r

	rr := *(*iface)(unsafe.Pointer(&r))
	ii := *(*iface)(unsafe.Pointer(&i))

	fmt.Printf("%v %v\n", ii.Data == rr.Data, ii.Data == unsafe.Pointer(b))
	fmt.Printf("%+v %+v\n", ii, rr)

	var i2 interface{} = &r
	ii2 := *(*iface)(unsafe.Pointer(&i2))
	fmt.Printf("%v %v\n", ii2.Data == rr.Data, ii2.Data == unsafe.Pointer(&r))
	fmt.Printf("%+v %+v\n", ii2, rr)
}

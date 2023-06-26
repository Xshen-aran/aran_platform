package test

import (
	"fmt"
	"testing"
)

type Data struct {
	Name string
	Foo  string
}

var d = Data{
	Name: "hello",
	Foo:  "world",
}

var s = []int{1, 2, 3}

func Test_print(t *testing.T) {
	fmt.Println(d)
	change(&d)
	fmt.Println(d)
	fmt.Println(s)
	changeSlice(s)
	fmt.Println(s)
}

func change(m *Data) {
	m.Foo = "change"
}

func changeSlice(s []int) {
	s[2] = 100
}

package main

import "fmt"

type sample struct {
	a int
	b int
}

func (s sample) AA() {
	fmt.Println(s.a, s.b)
}

type sample2 interface {
	AA()
}

func main() {
	ss := sample{1, 2}

	var ss2 sample2

	ss2 = ss
	ss2.AA()
}

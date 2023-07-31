package main

import "fmt"

type Namer interface {
	dsName() string
}

type Person struct {
	ame string
}

func (a Person) dsName() string {
	return a.ame
}

func main() {
	a := Person{"truong"}
	fmt.Println(a.dsName())
}

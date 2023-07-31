package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type School struct {
	Class string
	Point float32
}

func (p Person) String() string {
	return fmt.Sprintf("name : %s, age : %d", p.Name, p.Age)
}

func (th School) String() string {
	return fmt.Sprintf("class : %s, point : %f", th.Class, th.Point)
}

func main() {
	p1 := Person{"Alice", 18}
	fmt.Println(p1)

	p2 := School{"12A3", float32(10)}
	fmt.Println(p2)
}

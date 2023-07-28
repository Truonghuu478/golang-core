package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

type User struct {
	name string
	age  int
}

func main() {
	var v = Vertex{2, 3}
	fmt.Println("Abs : ", v.Abs())
}

func (data []User) displayUsers() {
	for user := range data {
		fmt.Println("Abs : ", user.name)

	}
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

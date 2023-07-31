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

func (user User) returnUser() {
	fmt.Println("user age:", user.age)
}

func displayUsers(data []User) {
	for _, user := range data {
		user.returnUser()
	}
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

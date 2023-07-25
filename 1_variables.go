package main

import "fmt"

var globalName string = "Truong"

func main() {
	userName := "Vy"
	num1, num2 := 1, 2
	var (
		name  = "strung"
		age   = 12
		class = "12A"
	)
	const PI = 3.14
	fmt.Println(userName)
	fmt.Println(globalName)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(name, age, class)
	fmt.Println("PI", PI)
}

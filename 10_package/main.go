package main

import (
	"fmt"
	custom "gomain/custom"
	until "gomain/until"
)

func main() {
	defer fmt.Println(until.FormatString(12))
	custom.SayHello("Trường")
	fmt.Println(until.Math(2, 3))

}

package main

import "fmt"

func main() {
	name1 := "hello world"
	fmt.Printf(" s: %s\n", name1)

	//%s là byte
	//%c là char

	for i := 0; i < len(name1); i++ {
		fmt.Printf(" char : %c\n", name1[i])
	}
}

package main

import "fmt"

func main() {
	//if else

	if a := 2 * 3; a < 2 {
		fmt.Printf(" %d < 2", a)
	} else {
		fmt.Printf(" %d > 2", a)
	}

	switch b := 10; {
	case b < 5:
		fmt.Printf(" %d < 5", b)
	case b > 9:
		fmt.Printf(" %d < 9", b)
		fallthrough
	case b == 10:
		fmt.Printf(" %d > 10", b)

	default:
		fmt.Printf(" %d == 0", b)

	}

}

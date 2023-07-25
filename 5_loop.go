package main

import "fmt"

func main() {
	//loop for

	// for i := 0; i < 10; i++ {
	// 	defer fmt.Println(i)
	// }
	//forEach
	listNumber := [3]int{1, 2, 3}

	for i, v := range listNumber {
		if i == 2 {
			continue
		}
		fmt.Printf("i =  %d, v= %d \n", i, v)

	}

}

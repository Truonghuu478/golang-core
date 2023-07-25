package main

import "fmt"

func main() {
	//array
	var array = [2]int{2, 3}
	fruits := []string{"apple", "banana", "orange"}
	fruits = append(fruits, "coco")
	fmt.Println("array : ", array)
	fmt.Println("fruits : ", fruits)
	//slice
	newFruits := fruits[:3]
	fmt.Println("newFruits : ", newFruits)

}

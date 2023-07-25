package main

import "fmt"

func main() {
	array1 := []int{0, 1}
	array2 := []int{3, 4}

	array1 = append(array1, array2...)

	variadicArray(&array1, 4, 3, 2, 2, 54)

	for i, v := range array1 {
		fmt.Printf("  i = %d, v= %d \n", i, v)
	}
}

func variadicArray(array *[]int, data ...int) {
	*array = append(*array, data...)
}

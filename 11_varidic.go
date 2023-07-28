package main

import "fmt"

func main() {
	array1 := []int{1, 2, 3, 4}
	array2 := []int{5, 6, 7, 8}
	array3 := []int{}

	array1 = append(array1, array2...)
	sum(1, 2, 3, 4, 4, 5)
	fmt.Println(array1, len(array3), cap(array3))
	//cap là chỉ dung lượng
	//len là chiều dài
}

func sum(nums ...int) int {
	total := 0
	for num := range nums {
		total += num
	}
	return total
}

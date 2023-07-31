package main

import "fmt"

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Generic function
func sum[T Number](a, b T) T {
	return a + b
}

func volume[T Number](a, b T) T {
	return a * b
}

//func map qua từng phần tử trong slice và transform
func ()  {
	
}

func main() {
	
	
	fmt.Println("tich :",  volume(2, 3))
	fmt.Println("sum :",  sum(2, 3))

}

package main

import "fmt"

func main() {
	//Bool
	isTrue := true
	var isFalse bool = false
	fmt.Println(isTrue, isFalse)
	//&& ,||.!,==,!=
	//Number
	var numberN = 20
	//%T là lấy kiểu dữ liệu, %d là lấy dữ liệu
	fmt.Printf("value of numberN : %d, type of n: %T\n", numberN, numberN)
	//bitwise
	// a, b := 3, 10
	// 3:  0011 => 1*2^0 + 1*2^1
	// 10: 1010 => 0*2^0 + 1*2^1 + 0*2^2 + 1*2^3
	//Float
	var floatDemo float32 = 1.2
	fmt.Println(floatDemo)
	//String
	// fullName  := "truong"
	//Assertion
	a := 15
	b := 8.3
	sum := a + int(b)
	fmt.Printf("%T\n", sum)
}

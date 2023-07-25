package main

import (
	"fmt"
	"strconv"
)

func main() {
	calculateSalaryEmployee(2000000, 12)

	fmt.Println(demoHighOrderFunc(2, 3, testNhan))
}

// high order funtion
func demoHighOrderFunc(a, b int, callBack func(int, int) int) int {
	return callBack(a, b)
}

func testNhan(a, b int) int {
	return a * b
}

func calculateSalaryEmployee(price float64, time int) {
	fmt.Println("equal : ", strconv.FormatFloat(price*float64(time), 'f', -1, 64), "VND")
	//"f" : định dạng số thập phân
	//-1 là value mặc định
	//64 : là kiểu float64
}

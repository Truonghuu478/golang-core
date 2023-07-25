package main

import (
	"fmt"
	"strconv"
)

func main() {
	var equal = calculateSalaryEmployee(2000000, 12)
	fmt.Println("equal : ", strconv.FormatFloat(equal, 'f', -1, 64), "VND")
	//"f" : định dạng số thập phân
	//-1 là value mặc định
	//64 : là kiểu float64
}

func calculateSalaryEmployee(price float64, time int) float64 {
	return price * float64(time)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var x int = 42
	var ptr *int = &x // con trỏ ptr trỏ tới biến x
	fmt.Println(*ptr) // in ra giá trị của biến mà ptr đang trỏ tới (42)
	*ptr++
	fmt.Printf("value of b: %d\n", *ptr)

	d := new(string) // zero value của string: ""
	fmt.Printf("value: %s, type: %T, address: %v\n", *d, d, d)
	*d = "Hello"
	var nums = []float32{1, 2, 2, 2, 22, 3, 1}

	var dtb float64

	handleDTB(nums, &dtb)
	refunded := math.Round(dtb)
	fmt.Println("DTB :", refunded)

	var aaa = &dtb
	var bbb = &aaa
	fmt.Println("pZ:", aaa, "address:", &aaa, "value:", *aaa)
	fmt.Println("ppZ:", bbb, "address:", &bbb, "value:", **bbb)

}

func handleDTB(nums []float32, dtb *float64) {
	var sum float32 = 0.0
	for i := 0; i < len(nums); i++ {
		sum += float32(nums[i])
	}
	if len(nums) > 0 {
		*dtb = float64(sum / float32(len(nums)))
	} else {
		*dtb = 0
	}
}

package main

import "fmt"

type HocSinh struct {
	fullName int
	class    string
	DTB      float32
	xepLoai  string
}

func main() {
	/**
	var name sting = "truong"
	name1:="vy"
	var name1,name2 := "truong","Vy"
	*/
	ex1()
}

func ex1() {
	var num1, num2, tong, tich int

	fmt.Println("Nhập số thứ nhất :")
	fmt.Scan(&num1) //toán tử & để lấy địa chỉ của biến khi sử dụng hàm fmt.Scan(), để cho phép hàm gán giá trị nhập vào cho biến được khai báo trước đó.
	fmt.Println("Nhập số thứ hai : ")
	fmt.Scan(&num2)
	tong = num1 + num2
	tich = num1 * num2

	fmt.Println("Tổng của 2 số là :", tong)
	fmt.Println("Tich của 2 số là :", tich)
}

func managerSV() {
	/**
	them hoc sinh
	hienthi danh sach
	tim kiem
	xoa
	sua
	*/

}

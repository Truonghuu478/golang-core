package custom

import "fmt"

func SayHello(name string) {
	fmt.Println("Hello ", name)
}

// Hàm init không thể được gọi ở bất cứ đâu
func init() {
	fmt.Println("Custom installing...")
}

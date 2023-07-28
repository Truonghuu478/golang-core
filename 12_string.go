package main

import "fmt"

func main() {
	const fullName string = "hello world"
	// for i := 0; i < len(fullName); i++ {
	// 	fmt.Printf(" %c", fullName[i])
	// }

	//cut hello

	hello := fmt.Sprintf("%s", fullName[:6])

	fmt.Println(hello, fmt.Sprintf("%s", fullName[6:11]))

}

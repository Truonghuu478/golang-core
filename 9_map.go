package main

import "fmt"

func main() {
	mapNum := make(map[string]int)
	//add key
	// fmt.Println("mapNum:", mapNum)
	for i := 1; i < 10; i++ {
		key := fmt.Sprintf("%d", i)
		mapNum[key] = i
	}

	//truy cap vao map
	fmt.Println("mapNum:", mapNum)

	//remove map
	delete(mapNum, "1")
	// fmt.Println("mapNum 1:", mapNum["1"])
	//*note : không tồn tại mặc định value = int(0)

	//kt key có trong map hay không
	if number2, ok := mapNum["2"]; ok {
		fmt.Printf("has a number is : %v\n", number2)
	} else {
		fmt.Printf("not has a number is : %v\n", number2)
	}

	for key, value := range mapNum {
		fmt.Printf("mapNum[%s] = %d\n", key, value)
	}

}

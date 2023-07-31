package main

import (
	"fmt"
	"time"
)

type AppError struct {
	Message string
	Code    int
	Time    time.Time
}

// Error implements error.
func (*AppError) Error() string {
	panic("unimplemented")
}

type Person struct {
	fullName string
	age      int
}

func (p *Person) newPerson(fullName string, age int) (*Person, error) {
	p.age = age

	if fullName == "" || age < 1 {
		return p, fmt.Errorf("Validate error")
	} else {
		p.fullName = fullName
		p.age = age
		return p, nil
	}
}

func mathTich(number int) (int, error) {
	if number < 0 {
		//  return 0, fmt.Errorf("Không thể chia cho số 0")
		return 0, fmt.Errorf("Không được nhập số < 0")
	}
	return number * 12, nil
}

func main() {
	result, error := mathTich(12)

	if error != nil {
		fmt.Println(error)
		return
	}

	fmt.Println("result : ", result)
	//case 2
	person := Person{}

	_, error = person.newPerson("truong", 0)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println("Create successful !", person)

	}

}

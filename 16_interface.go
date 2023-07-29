package main

type Namer interface {
	dsName() string
}

type Person struct {
	ame string
}

func (a Person) dsName() string {
	return a.ame
}

func main() {

}

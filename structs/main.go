package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}
func (p *Person) changeName(name string) {
	p.Name = name
}
func main() {
	myPerson := Person{
		Name: "benni",
	}
	fmt.Println(myPerson)

	// format with %+v
	fmt.Printf("my person %+v\n", myPerson)

	anotherPerson := NewPerson("Bennii", 32)
	anotherPerson.changeName("BenniNew")

	fmt.Printf("func call for person: %+v\n", anotherPerson)

	a := 7
	b := &a
	fmt.Println(b)
	fmt.Println(*b)

	mySlice := []int{
		1, 2, 3,
	}

	for index, _ := range mySlice {
		mySlice[index]++
	}

	fmt.Println(mySlice)
}

package topics

import "fmt"

type Person struct {
	Name string
	Age  int
}

func StructExample() {
	fmt.Println("Struct Example")

	p1 := Person{Name: "John", Age: 30}
	fmt.Println("Person:", p1)

	var p2 Person = Person{"Alice", 25}
	fmt.Println("Person:", p2)

	p3 := struct {
		fname string
		lname string
	}{fname: "Bob", lname: "Smith"}
	fmt.Println("Anonymous Person:", p3)

	persons := make([]Person, 2)
	persons[0] = p1
	persons[1] = p2

	// iterating through slice of structs
	for _, person := range persons {
		fmt.Println("Person from slice:", person)
	}
}

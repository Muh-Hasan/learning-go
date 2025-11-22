package topics

import "fmt"

func PointerExample() {
	fmt.Println("Pointer Example")

	var a int = 42

	var p *int = &a // p is a pointer to an int, initialized to the address of a

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of p (address of a):", p)
	fmt.Println("Value pointed to by p:", *p) // dereferencing p to get the value of a

	*p = 100 // changing the value of a through the pointer p
	fmt.Println("New value of a:", a)
}

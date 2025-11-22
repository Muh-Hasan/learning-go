package topics

import "fmt"

func ArrayExample() {
	fmt.Println("Array Example")
	// Few examples of declearing and initializing array
	var arr1 [3]int                  // Declaration of an array of integers with size 3
	arr2 := [3]string{"A", "B", "C"} // Declaration and initialization of an array of strings

	// iterating through array
	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr2[i])
	}

	for index, value := range arr2 {
		fmt.Println(index, value)
	}
}

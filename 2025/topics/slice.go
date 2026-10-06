package topics

import "fmt"

func SliceExample() {
	fmt.Println("Slice Example")
	// Few examples of declaring and initializing slices
	var slice1 []int                                             // Declaration of a slice of integers
	fmt.Println(slice1, len(slice1), cap(slice1), slice1 == nil) // Point to nil

	// ERROR:  slice1[1] = 0 // This will cause runtime panic: index out of range

	var slice2 = make([]int, 0) // Declaration and initialization of a slice of integers with length 0

	fmt.Println(slice2, len(slice2), cap(slice2), slice2 == nil) // [ ]

	var slice3 = make([]int, 3) // Declaration and initialization of a slice of integers with length 3

	slice3[0] = 1
	slice3 = append(slice3, 2)

	fmt.Println(slice3) // [1 0 0 2]

	slice4 := slice3[:] // Slicing the slice
	//  c := make([]string, len(s))
	// copy(c, s)

	fmt.Println(slice4) // [1 0 0 2]

	slice5 := []string{"A", "B", "C"} // Declaration and initialization of a slice of strings

	// iterating through slice
	for index, value := range slice5 {
		fmt.Println(index, value)
	}
}

// func sum(numbers ...int) int {
// 	total := 0

// 	for _, number := range numbers {
// 		total += number
// 	}

// 	return total
// }

// sum(1, 2, 3, 4, 5) // returns 15
// sum([]int{1, 2, 3, 4, 5}...) // returns 15
// sum(nums...) // where nums is a slice of integers

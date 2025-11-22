package topics

import (
	"fmt"
	"maps"
)

func MapExample() {
	fmt.Println("Map Example")
	// Few examples of declaring and initializing maps

	var map1 map[string]int        // Declaration of a map with string keys and int values
	fmt.Println(map1, map1 == nil) // true

	map1 = make(map[string]int) // Initializing the map

	fmt.Println(map1, map1 == nil) // false

	map2 := map[string]string{ // Declaration and initialization of a map with string keys and string values
		"A": "Apple",
		"B": "Banana",
		"C": "Cherry",
	}

	map2["D"] = "Date" // Adding a new key-value pair to the map

	fmt.Println(map2["D"])
	fmt.Println(map2)

	// iterating through map
	for key, value := range map2 {
		fmt.Println(key, value)
	}

	if value, exists := map2["A"]; exists {
		fmt.Println("Value:", value, exists)
	}

	if maps.Equal(map1, map1) {
		fmt.Println("map1 is equal to map1")
	}

	delete(map2, "D") // Deleting a key-value pair from the map

	fmt.Println(map2)

	clear(map2)

	fmt.Println("After clearing:", map2)
}

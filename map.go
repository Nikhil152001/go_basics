package main

import "fmt"

func main() {

	myMap := map[string]int{
		"nikhil":  5,
		"friend": 7,
		"sonal": 3,
	}

	key := "friend"

	if value, ok := myMap[key]; ok {

		fmt.Printf("Key '%s' is present in the map with value %d\n", key, value)
	} else {

		fmt.Printf("Key '%s' is NOT present in the map\n", key)
	}
}

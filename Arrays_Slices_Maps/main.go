package main

import (
	"fmt"
)

func main() {
	var ara [5] int
	ara[0] = 1
	ara[1] = 3
	ara[4] = 5
	fmt.Println(ara)
	fmt.Println(ara[0])
	
	for i := 0; i < len(ara); i++ {
		fmt.Println(ara[i])
	}
	
	for index, value := range ara {
		fmt.Println(index, value)
	}
	
	// say we want to iterate like this, but don't wanna use index
	// so, to avoid getting compile time error, we use the '_' underscore, it can't be used as index
	for _, value := range ara {
		fmt.Println(value)
	}
	
	// assign these values in first 5 indices & rests are filled with zeros
	// Notice: we can use an extra comma at the last position, mainly to easily comment out some values from last
	x := [10] int {98, 93, 77, 82, 83, }
	fmt.Println(x)
	
	// we can be break this down into separate lines
	y := [5] string {
		"gone",
		"with",
		"the",
		"wind",
	}
	
	fmt.Println(y)
	
	// no need to define size of the array with three dotes
	numbers := [...] int {1, 2, 3, 4, 5}
	fmt.Println(numbers)
}

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
}

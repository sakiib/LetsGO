package main

import (
	"fmt"
)

func main() {
	n := 10
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
	
	for i := 1; i <= n; i++ {
		if i % 2 == 0 {
			fmt.Printf("%v is Even\n", i)
		} else {
			fmt.Printf("%v is Odd\n", i)
		}
	}
	
	for i := 1; i <= n; i++ {
		reminder := i % 2
		switch reminder {
			case 0: fmt.Println("Even")
			case 1: fmt.Println("Odd")
		}
	}
}
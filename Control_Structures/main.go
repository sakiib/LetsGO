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
	
	m := make(map[string] int)
	// we can improve the conditionals also, there's a semicolon OMG!
	if val, ok := m["some_key"]; ok {
		fmt.Println(val)
	}
}

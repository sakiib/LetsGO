package main

import (
	"fmt"
)

func main() {
	stack := make([] int, 0)
	stack = append(stack, 1)
	stack = append(stack, 2)
	stack = append(stack, 3)
	fmt.Println(stack)
	for len(stack) > 0 {
	    	n := len(stack) - 1
        	fmt.Println(stack[n])
        	stack = stack[:n] 
	}
}

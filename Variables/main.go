package main

import (
	"fmt"
)

// package level or global declaration? Be careful, shorthand declaration can't be used here.
var num int = 100

func main() {
	
	// this is a comment 
	
	// declaring variables in different ways
	var i int 
	i = 5
	
	// declaring a variable j which is an int type variable
	var j int = 10
	
	// shorthand for declarign & assigning values to a new variable
	k:= 15
	
	// Printf - value type newline
	// should print: 5, int
	fmt.Printf("%v %T\n", i, i)
	
	// Print without & with a new line, takes one or more comma separated arguments
	fmt.Print(i, j, k)
	fmt.Println(i, j, k)
	
}

package main

import (
	"fmt"
)


// instead of declaring var everytime like this.. 
var firstName string = "sakib"
var lastName string = "alamin"
var company string = "appscode"
var id int = 12345

// .. we can do this.
var (
	firstName string = "sakib"
	lastName string = "alamin"
	company string = "appscode"
	id int = 12345
)

// package level or global declaration? Be careful, shorthand declaration can't be used here.
var num int = 100

// cannot declare variable multiple times under the same scope
// declared variables must be used somewhere in the code, compile-time error otherwise.

func main() {
	
	// this is a comment 
	
	// declaring variables in different ways
	var i int 
	i = 5
	
	// declaring a variable j which is an int type variable
	var j int = 10
	
	// if we don't mention the type, then it'll see the value & assign the type, string for this instance
	var name = "sakib"
	
	// shorthand for declarign & assigning values to a new variable
	k:= 15
	
	// Printf - value type newline
	// should print: 5, int
	fmt.Printf("%v %T\n", i, i)
	
	// Print without & with a new line, takes one or more comma separated arguments
	fmt.Print(i, j, k)
	fmt.Println(i, j, k)
	
	// MAIN DATA TYPES
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128
	
	var isPrime bool = false
	fmt.Printf("%v, %T\n", isPrime, isPrime)
	
	const number int = 10
	// can't redefine the value as number is a constant
	// number += 10
	
	fmt.Println(number + 10)
	
}

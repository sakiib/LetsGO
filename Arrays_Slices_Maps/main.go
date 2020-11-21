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
	
	// 2-D array
	var identityMatrix [3][3] int
	identityMatrix[0] = [3] int {1, 0, 0}
	identityMatrix[1] = [3] int {0, 1, 0}
	identityMatrix[2] = [3] int {0, 0, 1}
	fmt.Println(identityMatrix)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(identityMatrix[i][j], " ")
		}
		fmt.Println()
	}
	
	// array is just like values, doesn't point to the originl array while assigning, it just copies it
	a := [...] int {1, 2, 3}
	b := a
	b[1] = 10
	fmt.Println(a) // 1, 2, 3
	fmt.Println(b) // 1, 10, 3
	
	// slices, size need not to be defined
	// unlike array, slices actually points to the array a, from which b is assigned, so both go changed, like pointers do
	a := [] int {1, 2, 3}
	b := a
	b[1] = 10
	fmt.Println(a) // 1, 10, 3
	fmt.Println(b) // 1, 10, 3
	
	// length & capacity -> they aren't same
	fmt.Println(len(a))
	fmt.Println(cap(a))
	
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:] // slice of all elements
	c := a[3:] // slice from 4th element to end
	d := a[:6] // slice first 6 elements
	e := a[3:6] // slice the 4th, 5th, and 6th elements
	// they all point to same underlying data
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	
	// creating slice using make
	a := make([]int, 3)
	fmt.Println(a, len(a), cap(a)) // [0 0 0] 3 3
	// but, if we pass a third argument, it's the capacity of the slice, len & size need not to be the same
	a := make([]int, 3, 10)
	fmt.Println(a, len(a), cap(a)) // [0 0 0] 3 10
	
	
	// maps
	a := 1
	b := 1
	c := 3
	m := make(map[int] int)
	m[a] ++
	m[b] ++
	m[c] ++
	f := m[1]
	fmt.Println(f)
	delete(m, c)
	for key, val := range m {
	    fmt.Println(key, val)
	}
	
	m := make(map[string]int)
	// if key exists only, ok = true
	val, ok := m["some_key"]
	if ok {
		fmt.Println(val)
	}
	
	// if we need to check the existence only
	_, ok := m["some_key"]
	if ok {
		fmt.Println("key exists")
	}
	
	// we can improve the conditionals also, there's a semicolon OMG!
	if val, ok := m["some_key"]; ok {
		fmt.Println(val)
	}
}

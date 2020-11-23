package main

import (
	"fmt"
	"sort"
)

func main() {
	val := []int{3, 2, 1}
	sort.Ints(val)
	fmt.Println(val)
	
	// to sort an array []a, use it like a slice [l, r) -> r is exclusive, array indices 1 to n
	sort.Ints(a[1 : n+1])
}

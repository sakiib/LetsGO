package main

import (
	"fmt"
	"sort"
)

func main() {
	val := []int{3, 2, 1}
	sort.Ints(val)
	fmt.Println(val)
}

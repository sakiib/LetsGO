package main

import (
	"fmt"
)

func abs(x int, y int) int {
	if x >= y {
		return x - y
	} else {
		return y - x
	}
}

// using slices
func main() {
	var (
		n int
		d int
	)
	fmt.Scan(&n, &d)

	var ara []int
	for i := 0; i < n; i++ {
		var val int
		fmt.Scan(&val)

		ara = append(ara, val)
	}

	var ans int = 0
	for i, x := range ara {
		for j, y := range ara {
			if i != j && abs(x, y) <= d {
				ans++
			}
		}
	}

	fmt.Println(ans)
}

// using array
func main() {
	const N int = 1000 + 1
	var (
		n int
		d int
	)
	fmt.Scan(&n, &d)

	var ara [N]int
	for i := 0; i < n; i++ {
		fmt.Scan(&ara[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if abs(ara[i], ara[j]) <= d {
				ans++
			}
		}
	}

	fmt.Println(2 * ans)
}


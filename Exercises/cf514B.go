package main

import (
	"bufio"
	"fmt"
	"os"
)

type slope struct {
	up   int
	down int
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	have := make(map[slope]int)

	var n, X, Y int
	fmt.Fscanf(in, "%d %d %d\n", &n, &X, &Y)

	for i := 1; i <= n; i++ {
		var x, y int
		fmt.Fscanf(in, "%d %d\n", &x, &y)
		up, down := Y-y, X-x
		g := gcd(up, down)
		up /= g
		down /= g

		have[slope{up: up, down: down}]++
	}

	fmt.Println(len(have))
}

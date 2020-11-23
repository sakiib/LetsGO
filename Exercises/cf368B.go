package main

import (
	"bufio"
	"fmt"
	"os"
)

const N int = 1e5 + 1

var (
	a   [N]int
	ans [N]int
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var n, q int
	fmt.Fscanf(in, "%d %d\n", &n, &q)

	for i := 1; i <= n; i++ {
		fmt.Fscanf(in, "%d ", &a[i])
	}
	fmt.Fscanf(in, "\n")

	have := make(map[int]int)
	for i := n; i >= 1; i-- {
		have[a[i]]++
		ans[i] = len(have)
	}

	for qq := 1; qq <= q; qq++ {
		var idx int
		fmt.Fscanf(in, "%d\n", &idx)
		fmt.Println(ans[idx])
	}
}

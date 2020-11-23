package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)

	var (
		n int64
		k int
	)
	fmt.Fscanf(in, "%d %d\n", &n, &k)

	getDivisors := func() []int64 {
		var ret []int64
		for i := 1; int64(i)*int64(i) <= n; i++ {
			val := int64(i)
			if n%val == 0 {
				ret = append(ret, val)
				if n/val != val {
					ret = append(ret, n/val)
				}
			}
		}
		sort.SliceStable(ret, func(i, j int) bool {
			return ret[i] < ret[j]
		})
		return ret
	}

	divisors := getDivisors()

	if k > len(divisors) {
		fmt.Printf("-1\n")
	} else {
		fmt.Printf("%d\n", divisors[k-1])
	}
}

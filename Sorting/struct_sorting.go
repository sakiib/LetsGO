package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type data struct {
	x int
	y int
}

func main() {
	in := bufio.NewReader(os.Stdin)

	var strength, dragons int
	fmt.Fscanf(in, "%d %d\n", &strength, &dragons)

	var list []data
	for i := 1; i <= dragons; i++ {
		var x, y int
		fmt.Fscanf(in, "%d %d\n", &x, &y)
		list = append(list, data{x: x, y: y})
	}

	sort.SliceStable(list, func(i, j int) bool {
		if list[i].x == list[j].x {
			return list[i].y > list[j].y
		} else {
			return list[i].x < list[j].x
		}
	})

	ok := true
	for _, cur := range list {
		if strength <= cur.x {
			ok = false
			break
		} else {
			strength += cur.y
		}
	}

	if ok {
		fmt.Printf("YES\n")
	} else {
		fmt.Printf("NO\n")
	}
}

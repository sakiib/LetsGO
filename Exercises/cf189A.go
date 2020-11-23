package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
var n int
var a [5]int
var dp [4005]int
 
func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
 
func solve(rem int) int {
	if rem == 0 {
		return 0
	}
	if dp[rem] != -1 {
		return dp[rem]
	}
	ret := -1000000000
 
	for i := 1; i <= 3; i++ {
		if a[i] <= rem {
			ret = max(ret, solve(rem-a[i])+1)
		}
	}
 
	dp[rem] = ret
	return dp[rem]
}
 
func main() {
	in := bufio.NewReader(os.Stdin)
 
	fmt.Fscanf(in, "%d %d %d %d\n", &n, &a[1], &a[2], &a[3])
 
	for i := 0; i <= 4000; i++ {
		dp[i] = -1
	}
 
	fmt.Printf("%d\n", solve(n))
}

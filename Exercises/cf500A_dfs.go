package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
const N int = 1e5 + 5
 
var n, k int
var a [N]int
var graph [N][]int
 
func dfs(cur int) bool {
	if cur == k {
		return true
	}
	if cur >= n {
		return false
	}
 
	if dfs(cur + a[cur]) {
		return true
	} else {
		return false
	}
}
 
func main() {
	in := bufio.NewReader(os.Stdin)
 
	fmt.Fscanf(in, "%d %d\n", &n, &k)
 
	for i := 1; i < n; i++ {
		fmt.Fscanf(in, "%d ", &a[i])
		graph[i] = append(graph[i], i+a[i])
	}
 
	if dfs(1) {
		fmt.Printf("YES\n")
	} else {
		fmt.Printf("NO\n")
	}
}

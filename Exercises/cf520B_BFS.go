package main
 
import (
	"fmt"
)
 
const MAX int = 2e4
 
func bfs(n, m int) int {
    queue := make([]int, 0)
    queue = append(queue, n)
 
    var dist [MAX + 1] int
    for i := 0; i <= MAX; i++ {
        dist[i] = -1
    }
    
    dist[n] = 0
    
    for len(queue) > 0 {
        f := queue[0] 
        queue = queue[1:] 
        
        if f * 2 <= MAX && dist[f * 2] == -1 {
            dist[f * 2] = dist[f] + 1
            queue = append(queue, f * 2)
        }
        
        if f - 1 >= 1 && dist[f - 1] == -1 {
            dist[f - 1] = dist[f] + 1
            queue = append(queue, f - 1)
        }
    }
    return dist[m]
}
 
func main() {
	var n, m int
	fmt.Scan(&n, &m)
	fmt.Println(bfs(n, m))
}

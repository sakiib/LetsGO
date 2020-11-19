package main
 
import (
	"fmt"
)
 
func main() {
  const N int = 1000 + 1
  var ara [N][3]int 
    
	var problems int
	fmt.Scan(&problems)
	for i := 0; i < problems; i++ {
	    for j := 0; j < 3; j++ {
	        fmt.Scan(&ara[i][j])
	        if j > 0 {
	            ara[i][j] += ara[i][j - 1]
	        }
	    }
	}
	
	ans := 0
	for i := 0; i < problems; i++ {
	    if ara[i][2] >= 2 {
	        ans ++
	    }
	}
	
	fmt.Println(ans)
}

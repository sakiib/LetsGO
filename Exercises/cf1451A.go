package main
 
import (
	"fmt"
)
 
func main() {
	var t int
	fmt.Scan(&t)
 
	for tc := 1; tc <= t; tc++ {
		var n int
		fmt.Scan(&n)
		if n <= 3 {
			fmt.Println(n - 1)
		} else {
			if (n-3)%2 == 0 {
				fmt.Println(3)
			} else {
				fmt.Println(2)
			}
		}
	}
}

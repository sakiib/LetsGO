package main
 
import (
	"fmt"
)
 
const N int = 1e2 + 5
 
func main() {
	var t int
	fmt.Scan(&t)
 
	for tc := 1; tc <= t; tc++ {
		var length, query int
		var str string
		fmt.Scan(&length, &query, &str)
 
		str = "#" + str
		length = len(str)
 
		var prefixSum [N][2]int
		for i := 1; i < length; i++ {
			for j := 0; j <= 1; j++ {
				if int(str[i]-'0') == j {
					prefixSum[i][j]++
				}
				prefixSum[i][j] += prefixSum[i-1][j]
			}
		}
 
		ok := func(l int, r int, id int) bool {
			if l > r || l < 1 || r >= length {
				return false
			}
			if (prefixSum[r][id] - prefixSum[l-1][id]) > 0 {
				return true
			} else {
				return false
			}
		}
 
		for query > 0 {
			var l, r int
			fmt.Scan(&l, &r)
			if ok(1, l-1, int(str[l]-'0')) || ok(r+1, length-1, int(str[r]-'0')) {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
			query--
		}
	}
}

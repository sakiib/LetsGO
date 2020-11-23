package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
const N int = 1e6 + 5
 
type Info struct {
	bracket string
	index   int
}
 
func main() {
	in := bufio.NewReader(os.Stdin)
 
	var str string
	fmt.Fscanf(in, "%s\n", &str)
	str = "#" + str
 
	stack := make([]Info, 0)
	var dp [N]int
 
	length := len(str)
	for i := 1; i < length; i++ {
		cur := str[i : i+1]
		if cur == "(" {
			stack = append(stack, Info{bracket: "(", index: i})
		} else {
			if len(stack) == 0 {
				continue
			} else {
				n := len(stack) - 1
				top := stack[n]
				if top.bracket == "(" {
					dp[i] = i - top.index + 1 + dp[top.index-1]
					stack = stack[:n]
				} else {
					stack = append(stack, Info{bracket: ")", index: i})
				}
			}
		}
	}
 
	mx, cnt := 0, 0
	for i := 1; i < length; i++ {
		if dp[i] > mx {
			mx = dp[i]
		}
	}
	for i := 1; i < length; i++ {
		if dp[i] == mx {
			cnt++
		}
	}
	if mx == 0 {
		cnt = 1
	}
 
	fmt.Printf("%d %d\n", mx, cnt)
}

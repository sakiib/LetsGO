package main
 
import (
	"fmt"
)
 
const limit int = 1e2 + 10
var luckyNumbers []int
 
func generateLuckyNumbers(value int) {
    if value > limit {
        return
    }
    if value > 0 {
        luckyNumbers = append(luckyNumbers, value)
    }
    generateLuckyNumbers(value * 10 + 4)
    generateLuckyNumbers(value * 10 + 7)
}
 
func main() {
    generateLuckyNumbers(0)
    var n int64
    fmt.Scan(&n)
    
    cnt := 0
    for n > 0 {
        if n % 10 == 4 || n % 10 == 7 {
            cnt ++
        }
        n /= 10
    }
    
    ok := false
    for _, val := range luckyNumbers {
        if val == cnt {
            ok = true
        }
    }
    
    if ok {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

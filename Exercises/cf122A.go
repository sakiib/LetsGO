package main
 
import (
	"fmt"
)
 
const limit = 1000 + 10
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
    var n int
    fmt.Scan(&n)
    
    almostLucky := false
    for _, val := range luckyNumbers {
        if n % val == 0 {
            almostLucky = true
        }
    }
    
    if almostLucky {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

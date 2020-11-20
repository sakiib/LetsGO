package main

import (
	"fmt"
	"sort"
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

func binarySearch(cnt int) bool {
    lo := 0
    hi := len(luckyNumbers) - 1
    for lo <= hi {
        mid := (lo + hi) / 2
        if luckyNumbers[mid] == cnt {
            return true
        } else if luckyNumbers[mid] > cnt {
            hi = mid - 1
        } else {
            lo = mid + 1
        }
    }
    return false
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
    
    sort.Ints(luckyNumbers)
    //fmt.Println(luckyNumbers)
    
    ok := binarySearch(cnt)
    
    if ok {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

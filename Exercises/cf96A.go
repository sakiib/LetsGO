package main

import (
	"fmt"
)

func main() {
    var str string
    fmt.Scan(&str)
    
    danger := false
    for i := 0; i < len(str); i += 0 {
        j := i
        for j < len(str) && str[i] == str[j] {
            j ++
        }
        if j - i >= 7 {
            danger = true
            break
        } else {
            i = j
        }
    }
    
    if danger {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

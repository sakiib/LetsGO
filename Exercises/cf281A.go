package main

import (
	"fmt"
	"strings"
)

func main() {
    var str string
    fmt.Scan(&str)
    for i := 0; i < len(str); i++ {
        cur := str[i : i + 1]
        if i == 0 {
            cur = strings.ToUpper(cur)
            fmt.Print(cur)
        } else {
            fmt.Print(cur)
        }
    }
    fmt.Println()
}

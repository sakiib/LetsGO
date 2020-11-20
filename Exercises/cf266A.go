package main

import (
	"fmt"
)

func main() {
    var (
        len int
        str string
    )
    fmt.Scan(&len, &str)
    
    ans := 0
    for i := 0; i < len; i += 0 {
        j := i
        for j < len && str[i] == str[j] {
            j ++
        }
        ans += (j - i - 1)
        i = j
    }
    
    fmt.Println(ans)
}

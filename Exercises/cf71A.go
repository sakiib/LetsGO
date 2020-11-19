package main
 
import (
	"fmt"
	"strconv"
)
 
func main() {
	var testCase int
	fmt.Scan(&testCase)
	for tc := 1; tc <= testCase; tc++ {
	    var word string
	    fmt.Scan(&word)
	    if len(word) > 10 {
	        first := word[0 : 1]
	        last := word[len(word) - 1 : len(word)]
	        fmt.Println(first + strconv.Itoa(len(word) - 2) + last)
	    } else {
	        fmt.Println(word)
	    }
	}
}

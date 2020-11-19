package main
 
import (
	"fmt"
)
 
func main() {
	var weight int
	fmt.Scan(&weight)
	if weight >= 4 && (weight - 2) % 2 == 0 {
	    fmt.Println("Yes")
	} else {
	    fmt.Println("No")
	}
}

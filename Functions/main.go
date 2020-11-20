package main
 
import (
	"fmt"
	"strings"
)

// func function_name (argument) return type
func isVowel(ch string) bool {
    if ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u" || ch == "y" {
        return true
    } else {
        return false
    }
}

var luckyNumbers []int

// recursive function to generate numbers with digits 4,7 only (lucky numbers)
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

// simple function to check if a number is present in the slices of lucky numbers
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

// function can return multiple values
func myFunc(x int, y int) (int, int) {
	return y, x
}

// variadic function
// By using ... before the type name of the last parameter we can indicate that it takes zero or more of those parameters.
// This is precisely how the fmt.Println function is implemented
 
func calcSum(name string, args ...int) int {
	sum := 0
	for _, val := range args {
		sum += val
	}
	return sum
}

func main() {
	first, second := myFunc(10, 20)
	fmt.Println(first, second)
	
	fmt.Println(calcSum("sakib", 1, 2, 3, 4, 5))
	
	// closures 
	// It is possible to create functions inside of functions
	// addTwo is a local variable that has the type func(int, int) int (a function that takes two ints and returns an int).
	addTwo := func(x, y int) int {
		return x + y
	}
	fmt.Println(addTwo(5, 10))
}

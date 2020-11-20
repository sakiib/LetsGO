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

func main() {
	first, second := myFunc(10, 20)
	fmt.Println(first, second)
}

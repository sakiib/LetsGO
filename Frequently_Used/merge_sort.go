package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
const N int = 2e5 + 5
 
var a [N]int
var ara [N]int
var temp [N]int
 
func merge(lo, mid, hi int) {
	k, left, right := lo, lo, mid+1
	for k <= hi {
		if left > mid {
			temp[k] = ara[right]
			k++
			right++
		} else if right > hi {
			temp[k] = ara[left]
			k++
			left++
		} else if ara[left] <= ara[right] {
			temp[k] = ara[left]
			k++
			left++
		} else {
			temp[k] = ara[right]
			k++
			right++
		}
	}
	for i := lo; i <= hi; i++ {
		ara[i] = temp[i]
	}
}
 
func merge_sort(lo, hi int) {
	if lo == hi {
		return
	}
	mid := (lo + hi) / 2
	merge_sort(lo, mid)
	merge_sort(mid+1, hi)
	merge(lo, mid, hi)
}
 
func main() {
	
	for i := 1; i <= n; i++ {
		ara[i] = a[i]
	}
	merge_sort(1, n)
	for i := 1; i <= n; i++ {
		a[i] = ara[i]
	}
 
}

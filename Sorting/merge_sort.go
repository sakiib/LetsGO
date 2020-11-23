package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
const N int = 2e5 + 5
 
var a [N]int
var ara [N]int
var impact [N]int
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
	in := bufio.NewReader(os.Stdin)
 
	var n, q int
	fmt.Fscanf(in, "%d %d\n", &n, &q)
	for i := 1; i <= n; i++ {
		fmt.Fscanf(in, "%d ", &a[i])
	}
	fmt.Fscanf(in, "\n")
 
	for qq := 1; qq <= q; qq++ {
		var l, r int
		fmt.Fscanf(in, "%d %d\n", &l, &r)
		impact[l]++
		impact[r+1]--
	}
	for i := 1; i <= n; i++ {
		impact[i] += impact[i-1]
	}
 
	for i := 1; i <= n; i++ {
		ara[i] = a[i]
	}
	merge_sort(1, n)
	for i := 1; i <= n; i++ {
		a[i] = ara[i]
	}
 
	for i := 1; i <= n; i++ {
		ara[i] = impact[i]
	}
	merge_sort(1, n)
	for i := 1; i <= n; i++ {
		impact[i] = ara[i]
	}
 
	var ans int64 = 0
	for i := 1; i <= n; i++ {
		ans += int64(a[i]) * int64(impact[i])
	}
 
	fmt.Printf("%d\n", ans)
}

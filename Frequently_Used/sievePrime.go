package main

import (
	"fmt"
)

const N int = 1e5 + 5

var isPrime [N]bool

func sievePrime(n int) []int {
	for i := 0; i <= n; i++ {
		isPrime[i] = true
	}
	isPrime[0] = false
	isPrime[1] = false
	for i := 4; i <= n; i += 2 {
		isPrime[i] = false
	}
	for i := 3; i*i <= n; i += 2 {
		if isPrime[i] {
			for j := i * i; j <= n; j += 2 * i {
				isPrime[j] = false
			}
		}
	}
	var prime []int
	prime = append(prime, 2)
	for i := 3; i <= n; i += 2 {
		if isPrime[i] {
			prime = append(prime, i)
		}
	}
	return prime
}

func main() {
	prime := sievePrime(100)
	for _, val := range prime {
		fmt.Println(val)
	}
}

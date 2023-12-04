package main

import (
	"fmt"

	"github.com/tebeka/deque"
)

var N, K = 41, 3

// Josephus problem
// https://en.wikipedia.org/wiki/Josephus_problem

// Iterative solution using modulus
func modulus(n int, k int) int {
	var result int
	for i := 2; i <= n; i++ {
		result = (result + k) % i
	}
	return result + 1
}

// Double ended queue solution
func dequeue(n int, k int) int {
	dq := deque.New()
	for i := 1; i <= n; i++ {
		dq.Append(i)
	}
	count := 0
	for dq.Len() > 1 {
		count++
		if count%k == 0 {
			dq.PopLeft()
		} else {
			elem, _ := dq.PopLeft()
			dq.Append(elem)
		}
	}
	res, _ := dq.PopLeft()
	return res.(int)
}

func main() {
	N, K = 10_000_000, 67
	fmt.Printf("The survivor according to modulus is %d\n", modulus(N, K))
	fmt.Printf("The survivor according to deque is %d\n", dequeue(N, K))
}

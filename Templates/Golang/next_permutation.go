package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	wt = bufio.NewWriter(os.Stdout)
)

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func nextPermutation(a []int) bool {
	n := len(a)
	reverse := func(begin, end int) {
		for i := 0; i < (end-begin)/2; i++ {
			a[begin+i], a[end-1-i] = a[end-1-i], a[begin+i]
		}
	}
	for i := n - 2; i >= 0; i-- {
		if a[i] < a[i+1] {
			var j int
			for j = n - 1; a[i] >= a[j]; j-- {
			}
			a[i], a[j] = a[j], a[i]
			reverse(i+1, n)
			return true
		}
	}
	return false
}

func main() {
	defer wt.Flush()

	a := []int{1, 2, 3, 4, 5}

	puts(a)
	for nextPermutation(a) {
		puts(a)
	}
}

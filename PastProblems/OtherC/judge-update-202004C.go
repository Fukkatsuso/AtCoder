package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

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

func sum(a []int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a := nextInts(3)

	n := sum(a)
	b := make([]int, n)
	for i := range b {
		b[i] = i + 1
	}

	ans := 0
	for {
		ok := true
		bi := 0
		for i := 0; i < 3; i++ {
			for j := 0; j < a[i]; j, bi = j+1, bi+1 {
				// 左のブロックより大きいか
				if i > 0 && b[bi] > b[bi-a[i-1]] {
					ok = false
				}
				// 下のブロックより大きいか
				if j > 0 && b[bi] > b[bi-1] {
					ok = false
				}
			}
		}
		if ok {
			ans++
		}
		if !nextPermutation(b) {
			break
		}
	}

	puts(ans)
}

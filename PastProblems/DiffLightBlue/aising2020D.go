// 解説AC

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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func popcount(x int) int {
	res := 0
	for x > 0 {
		res += (x & 1)
		x >>= 1
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, x := nextInt(), next()

	f := make([]int, 200001)
	for i := 1; i <= 200000; i++ {
		p := popcount(i)
		f[i] = f[i%p] + 1
	}
	pc := 0
	for i := range x {
		if x[i] == '1' {
			pc++
		}
	}
	if pc == 0 {
		for i := 0; i < n; i++ {
			puts(1)
		}
		return
	} else if pc == 1 {
		for i := 0; i < n; i++ {
			if x[i] == '1' {
				puts(0)
			} else if i == n-1 || x[n-1] == '1' {
				// xi%2 == 1となるとき
				puts(2)
			} else {
				puts(1)
			}
		}
		return
	}

	// r[i]: 2^(n-1-i)を[pc-1, pc+1]で割ったあまり
	r := make([][2]int, n)
	// xr: xを[pc-1, pc+1]で割ったあまり
	xr := [2]int{}
	pow0, pow1 := 1, 1
	for i := n - 1; i >= 0; i-- {
		r[i][0], r[i][1] = pow0, pow1
		if x[i] == '1' {
			xr[0] += r[i][0]
			xr[0] %= pc - 1
			xr[1] += r[i][1]
			xr[1] %= pc + 1
		}
		pow0 *= 2
		pow0 %= pc - 1
		pow1 *= 2
		pow1 %= pc + 1
	}

	for i := 0; i < n; i++ {
		var xi int
		if x[i] == '0' {
			// 0 -> 1
			xi = (xr[1] + r[i][1]) % (pc + 1)
		} else {
			// 1 -> 0
			xi = (xr[0] - r[i][0] + pc - 1) % (pc - 1)
		}
		ans := 1 + f[xi]
		puts(ans)
	}
}

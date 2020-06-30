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

func nextInt2() (int, int) {
	return nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, c := nextInt2()
	x, v := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], v[i] = nextInt(), nextInt()
	}

	sv1 := make([]int, n)
	sv1[0] = v[0] - x[0]
	for i := 1; i < n; i++ {
		sv1[i] = sv1[i-1] + v[i] - (x[i] - x[i-1])
	}
	sv2 := make([]int, n)
	sv2[n-1] = v[n-1] - (c - x[n-1])
	for i := n - 2; i >= 0; i-- {
		sv2[i] = sv2[i+1] + v[i] - (x[i+1] - x[i])
	}

	vmax1 := make([]int, n)
	vmax1[0] = 0
	for i := 1; i < n; i++ {
		if sv1[i] > sv1[vmax1[i-1]] {
			vmax1[i] = i
		} else {
			vmax1[i] = vmax1[i-1]
		}
	}
	vmax2 := make([]int, n)
	vmax2[n-1] = n - 1
	for i := n - 2; i >= 0; i-- {
		if sv2[i] > sv2[vmax2[i+1]] {
			vmax2[i] = i
		} else {
			vmax2[i] = vmax2[i+1]
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		get := sv1[vmax1[i]]
		if i < n-1 {
			getRev := get - x[vmax1[i]] + sv2[vmax2[i+1]]
			ans = max(ans, get, getRev)
		} else {
			ans = max(ans, get)
		}
	}
	for i := n - 1; i >= 0; i-- {
		get := sv2[vmax2[i]]
		if i > 0 {
			getRev := get - (c - x[vmax2[i]]) + sv1[vmax1[i-1]]
			ans = max(ans, get, getRev)
		} else {
			ans = max(ans, get)
		}
	}

	puts(ans)
}

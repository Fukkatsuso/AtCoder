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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func min(nums ...float64) float64 {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	t, v := make([]int, n+2), make([]int, n+2)
	t[0], v[0] = 0, 0
	for i := 1; i <= n; i++ {
		t[i] = getInt()
		t[i] += t[i-1]
	}
	for i := 1; i <= n; i++ {
		v[i] = getInt()
	}
	t[n+1], v[n+1] = t[n], 0

	// 時刻xで出せる最高速度
	vmax := func(x float64) float64 {
		res := 1e10
		for i := 0; i <= n+1; i++ {
			l, r := 0.0, 0.0
			if 0 < i {
				l, r = float64(t[i-1]), float64(t[i])
			}

			// 条件iを満たすための時刻xでの最高速度
			y := float64(v[i])
			if x < l {
				y += l - x
			} else if r < x {
				y += x - r
			}
			res = min(res, y)
		}
		return res
	}

	ans := 0.0
	for i := 0; i < 2*t[n]; i++ {
		a := float64(i) / 2.0
		d := (vmax(a) + vmax(a+0.5)) * 0.25
		ans += d
	}
	puts(ans)
}

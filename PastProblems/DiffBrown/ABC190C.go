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

	n, m := getInt(), getInt()
	a, b := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		a[i], b[i] = getInt()-1, getInt()-1
	}
	k := getInt()
	cd := make([][2]int, k)
	for i := 0; i < k; i++ {
		cd[i][0], cd[i][1] = getInt()-1, getInt()-1
	}

	ans := 0
	for bit := 0; bit < (1 << k); bit++ {
		// ボールを配置
		ball := make([]int, n)
		for i := 0; i < k; i++ {
			// 下からi番目のビットが立っていれば皿Diにボールを置く
			dish := cd[i][(bit>>i)&1]
			ball[dish]++
		}
		// 満たされる条件の個数
		count := 0
		for i := 0; i < m; i++ {
			if ball[a[i]] > 0 && ball[b[i]] > 0 {
				count++
			}
		}
		ans = max(ans, count)
	}

	puts(ans)
}

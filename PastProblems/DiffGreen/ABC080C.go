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

	n := nextInt()
	f, p := make([][10]int, n), make([][11]int, n)
	for i := range f {
		for j := range f[i] {
			f[i][j] = nextInt()
		}
	}
	for i := range p {
		for j := range p[i] {
			p[i][j] = nextInt()
		}
	}

	ans := -100000000 * n
	open := make([]int, 10)
	for bit := 1; bit < 1<<10; bit++ {
		// 営業するか否かを決定
		for hour := 0; hour < 10; hour++ {
			open[hour] = (bit >> uint(hour)) & 1
		}
		// 得られる利益を計算
		profit := 0
		for i := 0; i < n; i++ {
			// 各店舗と被っている時間帯数
			c := 0
			for j := 0; j < 10; j++ {
				c += open[j] * f[i][j]
			}
			profit += p[i][c]
		}
		// 最大値更新
		ans = max(ans, profit)
	}

	puts(ans)
}

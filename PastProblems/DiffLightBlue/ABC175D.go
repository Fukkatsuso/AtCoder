// 解説AC
// 終着点までの距離によって周回可能回数が変わることを考慮できなかった

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

func min(nums ...int) int {
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

	n, k := nextInt(), nextInt()
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = nextInt() - 1
	}
	c := nextInts(n)

	// マスfromからステップごとに得られるスコア
	calcPoint := func(from int) []int {
		point := make([]int, 0)
		sum := 0
		for i := from; ; i = p[i] {
			to := p[i]
			sum += c[to]
			point = append(point, sum)
			if to == from {
				break
			}
		}
		return point
	}

	ans := -10000000000
	for i := 0; i < n; i++ {
		// マスiから始めたときの最大スコアを求める
		point := calcPoint(i)
		size := len(point)
		var score int
		if k >= size && point[size-1] > 0 {
			// できるだけ多く周回
			// マスiから距離j+1の位置のマスを終着点としてスコア計算
			for j := range point {
				rap := (k - j - 1) / size
				score = max(score, point[size-1]*rap+point[j])
			}
		} else {
			// 周回はしないorできない
			score = max(point[:min(size, k)]...)
		}
		ans = max(ans, score)
	}

	puts(ans)
}

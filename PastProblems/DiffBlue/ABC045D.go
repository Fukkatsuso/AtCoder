// 解説AC
// 実装に詰まって参考程度に見た
// 考察はできていた

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

func sum(a ...int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

type Point struct {
	a, b int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w, n := getInt(), getInt(), getInt()
	mp := map[Point]bool{}
	for i := 0; i < n; i++ {
		p := Point{}
		p.a, p.b = getInt()-1, getInt()-1
		mp[p] = true
	}

	// マス(a,b)が黒か否か
	isBlack := func(a, b int) bool {
		return mp[Point{a, b}]
	}
	// マス(a,b)を中心にした領域の黒マスの数
	blackCount := func(a, b int) int {
		if a <= 0 || h-1 <= a || b <= 0 || w-1 <= b {
			return 0
		}
		res := 0
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				if isBlack(a+i, b+j) {
					res++
				}
			}
		}
		return res
	}

	ans := make([]int, 10)
	checked := map[Point]bool{}
	for p := range mp {
		// マス(p.a,p.b)から(i,j)だけずらした位置を中心にして，その領域の黒マスの数を数える
		// 重複(チェック済みの中心マス)は省く
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				center := Point{p.a + i, p.b + j}
				if !checked[center] {
					cnt := blackCount(p.a+i, p.b+j)
					ans[cnt]++
					checked[center] = true
				}
			}
		}
	}

	ans[0] = (h-2)*(w-2) - sum(ans[1:]...)
	for i := range ans {
		puts(ans[i])
	}
}

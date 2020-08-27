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

func max(nums []int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

type Point struct {
	h, w int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w, m := nextInt(), nextInt(), nextInt()

	sh, sw := make([]int, h), make([]int, w)
	set := map[Point]bool{}
	for i := 0; i < m; i++ {
		hi, wi := nextInt()-1, nextInt()-1
		sh[hi]++
		sw[wi]++
		set[Point{hi, wi}] = true
	}

	// 行，列ごとの最大の爆破数
	shmax, swmax := max(sh), max(sw)
	// 爆破対象に選ぶとshmax/swmaxだけ爆破できるような座標
	hmax, wmax := make([]int, 0), make([]int, 0)
	for i := 0; i < h; i++ {
		if shmax == sh[i] {
			hmax = append(hmax, i)
		}
	}
	for i := 0; i < w; i++ {
		if swmax == sw[i] {
			wmax = append(wmax, i)
		}
	}

	ans := shmax + swmax
	for i := range hmax {
		for j := range wmax {
			// 爆破数を最大化でき，爆弾と重ならないマスが存在すればOK
			if !set[Point{hmax[i], wmax[j]}] {
				puts(ans)
				return
			}
		}
	}
	// 爆弾と重なるマスしかない場合
	puts(ans - 1)
}

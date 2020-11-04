// 解説AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
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

	n := nextInt()
	b := make([]int, n)
	for i := range b {
		// 10^9をそのまま実数にかけると誤差が生じるので地道に変換
		a := next()
		split := strings.Split(a, ".")
		// 整数部
		x, _ := strconv.Atoi(split[0])
		b[i] = x * 1000000000
		// 小数部（あれば）
		if len(split) == 2 {
			y, _ := strconv.Atoi(split[1])
			for j := len(split[1]); j < 9; j++ {
				y *= 10
			}
			b[i] += y
		}
	}

	// tf[i][j]: b[x]を素因数分解したとき，(2^i)*(5^j)*...になる数
	tf := make([][19]int, 19)
	for i := range b {
		pf := map[int]int{}
		for b[i]%2 == 0 {
			pf[2]++
			b[i] /= 2
		}
		for b[i]%5 == 0 {
			pf[5]++
			b[i] /= 5
		}
		tf[min(pf[2], 18)][min(pf[5], 18)]++
	}

	ans := 0
	for t1 := 0; t1 <= 18; t1++ {
		for t2 := 0; t2 <= 18; t2++ {
			if t1+t2 < 18 {
				continue
			}
			for f1 := 0; f1 <= 18; f1++ {
				for f2 := 0; f2 <= 18; f2++ {
					if f1+f2 < 18 {
						continue
					}
					if t1 == t2 && f1 == f2 {
						ans += tf[t1][f1] * (tf[t1][f1] - 1)
					} else {
						ans += tf[t1][f1] * tf[t2][f2]
					}
				}
			}
		}
	}
	ans /= 2

	puts(ans)
}

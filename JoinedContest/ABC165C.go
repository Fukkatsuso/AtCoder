// TLE

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

func pow(a, n int) int {
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret *= a
		}
		a *= a
		n >>= 1
	}
	return ret
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

	_, m, q := nextInt(), nextInt(), nextInt()
	a, b, c, d := make([]int, q), make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		a[i], b[i], c[i], d[i] = nextInt(), nextInt(), nextInt(), nextInt()
	}

	min := func(mii map[int]int) int {
		res := 1000
		for _, v := range mii {
			if res > v {
				res = v
			}
		}
		return res
	}

	ans := 0
	// bit全探索
	iMax := 1 << uint(q)
	for i := 1; i < iMax; i++ {
		A := map[int]int{}
		ok := true
		point := 0
		// Aを構築
		for j := 0; j < q; j++ {
			if (i>>j)&1 == 0 {
				continue
			}
			// j番目のabcd組を選択
			Aaj, aok := A[a[j]]
			Abj, bok := A[b[j]]
			if aok && bok {
				if Abj-Aaj != c[j] {
					ok = false
					break
				}
			} else if aok {
				A[b[j]] = c[j] + Aaj
			} else if bok {
				A[a[j]] = Abj - c[j]
			} else {
				A[a[j]] = 0
				A[b[j]] = c[j]
			}
			point += d[j]
		}
		// Aが範囲内に収まるか確認
		minA := min(A)
		for _, v := range A {
			if v-minA+1 > m {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}
		// 得点計算
		ans = max(ans, point)
	}

	puts(ans)
}

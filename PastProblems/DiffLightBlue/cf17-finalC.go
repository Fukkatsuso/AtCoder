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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
	d := nextInts(n)

	cnt := make([]int, 13)
	for i := range d {
		cnt[d[i]]++
	}
	cnt[0]++

	// 0,12に複数人がいる場合
	if cnt[0] >= 2 || cnt[12] >= 2 {
		puts(0)
		return
	}

	// 割り振られた人数
	tm := make([]int, 24)
	mind := func() int {
		res := 24
		for i := 0; i < 24; i++ {
			if tm[i] == 0 {
				continue
			}
			for j := 0; j < 24; j++ {
				if tm[j] == 0 || i == j {
					continue
				}
				res = min(res, min(abs(i-j), 24-abs(i-j)))
			}
		}
		return res
	}

	// 割り振りが確定していないdi
	rem := make([]int, 0)

	// 同じ時間差地点に3人以上いれば必ず重複する
	for i := range cnt {
		if cnt[i] >= 3 {
			puts(0)
			return
		}
		if cnt[i] == 2 {
			tm[i], tm[24-i] = 1, 1
			cnt[i] = 0
		}
		if cnt[i] == 1 {
			rem = append(rem, i)
		}
	}
	if len(rem) == 0 {
		puts(mind())
		return
	}

	ans := 0
	m := len(rem)
	// bit全探索
	for i := 0; i < (1 << m); i++ {
		// 割り振り
		for j := range rem {
			if i&(1<<j) > 0 {
				tm[(24-rem[j])%24] = 1
			} else {
				tm[rem[j]] = 1
			}
		}
		ans = max(ans, mind())
		// 元に戻す
		for j := range rem {
			if i&(1<<j) > 0 {
				tm[(24-rem[j])%24] = 0
			} else {
				tm[rem[j]] = 0
			}
		}
	}
	puts(ans)
}

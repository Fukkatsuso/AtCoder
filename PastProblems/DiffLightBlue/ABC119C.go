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

func nextInt4() (int, int, int, int) {
	return nextInt(), nextInt(), nextInt(), nextInt()
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

// llの棒を一つにまとめて長さlに変えるためのコスト
func joinCost(ll []int, l int) int {
	if len(ll) == 0 {
		return 1 << 50
	}
	c, sum := 10*(len(ll)-1), 0
	for i := range ll {
		sum += ll[i]
	}
	return c + abs(l-sum)
}

// group: A, B, C, 使わない
func rec(l, group []int, n, id, a, b, c int) int {
	if id == n {
		ll := [][]int{[]int{}, []int{}, []int{}}
		for i := range group {
			if group[i] == 3 {
				continue
			}
			ll[group[i]] = append(ll[group[i]], l[i])
		}
		return joinCost(ll[0], a) + joinCost(ll[1], b) + joinCost(ll[2], c)
	}

	mp := 1 << 60
	for i := 0; i < 4; i++ {
		group[id] = i
		mp = min(mp, rec(l, group, n, id+1, a, b, c))
	}
	return mp
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, a, b, c := nextInt4()
	l := nextInts(n)

	puts(rec(l, make([]int, n), n, 0, a, b, c))
}

// 解説AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sum(n int, q [][]int) int {
	res := 0
	for i := 0; i < 2; i++ {
		for j := 1; j < len(q[i]); j++ {
			res += abs(q[i][j] - q[i][j-1])
		}
	}
	return res
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
	a := nextInts(n)
	sort.Sort(sort.IntSlice(a))

	// 大小の順
	q1 := make([][]int, 2)
	for i := range q1 {
		q1[i] = make([]int, 0)
	}
	q1[0], q1[1] = append(q1[0], a[0]), append(q1[1], a[0])
	for l, r := 1, n-1; l <= r; {
		q1[0] = append(q1[0], a[r])
		r--
		if l <= r {
			q1[1] = append(q1[1], a[r])
			r--
		}
		if l <= r {
			q1[0] = append(q1[0], a[l])
			l++
		}
		if l <= r {
			q1[1] = append(q1[1], a[l])
			l++
		}
	}

	// 小大の順
	q2 := make([][]int, 2)
	for i := range q2 {
		q2[i] = make([]int, 0)
	}
	q2[0], q2[1] = append(q2[0], a[n-1]), append(q2[1], a[n-1])
	for l, r := 0, n-2; l <= r; {
		q2[0] = append(q2[0], a[l])
		l++
		if l <= r {
			q2[1] = append(q2[1], a[l])
			l++
		}
		if l <= r {
			q2[0] = append(q2[0], a[r])
			r--
		}
		if l <= r {
			q2[1] = append(q2[1], a[r])
			r--
		}
	}

	puts(max(sum(n, q1), sum(n, q2)))
}

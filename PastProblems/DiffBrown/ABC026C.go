package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
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

func bfs(subs [][]int, salary []int, id int) int {
	if salary[id] > 0 {
		return salary[id]
	}
	if len(subs[id]) == 0 {
		salary[id] = 1
	} else if len(subs[id]) == 1 {
		salary[id] = 2*bfs(subs, salary, subs[id][0]) + 1
	} else {
		smax, smin := 0, 1<<60
		for _, sub := range subs[id] {
			sal := bfs(subs, salary, sub)
			smax = max(smax, sal)
			smin = min(smin, sal)
		}
		salary[id] = smax + smin + 1
	}
	return salary[id]
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	b := nextInts(n - 1)

	subs := make([][]int, n+1)
	for i := range subs {
		subs[i] = make([]int, 0)
	}
	for i := range b {
		subs[b[i]] = append(subs[b[i]], i+2)
	}

	salary := make([]int, n+1)
	for i := range salary {
		salary[i] = -1
	}

	puts(bfs(subs, salary, 1))
}

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

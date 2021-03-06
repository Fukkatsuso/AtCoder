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

type Emp struct {
	id, time int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	a := make([]Emp, n)
	b := make([]Emp, n)
	for i := 0; i < n; i++ {
		ai, bi := getInt(), getInt()
		a[i] = Emp{i, ai}
		b[i] = Emp{i, bi}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].time < a[j].time
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i].time < b[j].time
	})

	var ans int
	if a[0].id == b[0].id {
		ans = min(a[0].time+b[0].time, max(a[0].time, b[1].time), max(a[1].time, b[0].time))
	} else {
		ans = max(a[0].time, b[0].time)
	}
	puts(ans)
}

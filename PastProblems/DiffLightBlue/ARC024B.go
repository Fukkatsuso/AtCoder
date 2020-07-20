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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	col := nextInts(n)

	r := 0
	for i := range col {
		r += col[i]
	}
	if r == n || r == 0 {
		puts(-1)
		return
	}

	t := make([][2]int, 0)
	t = append(t, [2]int{col[0], 1})
	for i := 1; i < n; i++ {
		if col[i] == t[len(t)-1][0] {
			t[len(t)-1][1]++
		} else {
			t = append(t, [2]int{col[i], 1})
		}
	}
	if t[len(t)-1][0] == t[0][0] {
		u := t[len(t)-1][1]
		t = t[:len(t)-1]
		t[0][1] += u
	}

	ans := 1
	for i := range t {
		ans = max(ans, (t[i][1]+1)/2)
	}
	puts(ans)
}

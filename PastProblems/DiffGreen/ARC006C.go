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

func puttable(mount []int, w int) int {
	minDiff := 100000
	id := -1
	for i := range mount {
		if d := mount[i] - w; 0 <= d && d < minDiff {
			minDiff = d
			id = i
		}
	}
	return id
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	w := nextInts(n)

	mount := []int{w[0]}
	for i := 1; i < n; i++ {
		id := puttable(mount, w[i])
		if id < 0 {
			mount = append(mount, w[i])
		} else {
			mount[id] = w[i]
		}
	}

	puts(len(mount))
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Count struct {
	key, val int
}

type Counts []*Count

func (c Counts) Len() int           { return len(c) }
func (c Counts) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Counts) Less(i, j int) bool { return c[i].val > c[j].val }

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	m := make([]map[int]int, 2)
	for i := range m {
		m[i] = make(map[int]int, 0)
	}
	for i := 0; i < n; i++ {
		v := nextInt()
		m[i%2][v]++
	}

	cnts := make([]Counts, 2)
	for i := range cnts {
		cnts[i] = mapToCounts(m[i])
		sort.Sort(cnts[i])
	}

	var ans int
	if len(cnts[0])*len(cnts[1]) == 0 {
		ans = n / 2
	} else if cnts[0][0].key != cnts[1][0].key {
		ans = n - cnts[0][0].val - cnts[1][0].val
	} else {
		if len(cnts[0]) == 1 && len(cnts[1]) == 1 {
			ans = n / 2
		} else if len(cnts[0]) == 1 {
			ans = n/2 - cnts[1][1].val
		} else if len(cnts[1]) == 1 {
			ans = n/2 - cnts[0][1].val
		} else {
			ans = n - max(cnts[0][0].val+cnts[1][1].val, cnts[1][0].val+cnts[0][1].val)
		}
	}

	puts(ans)
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

func mapToCounts(m map[int]int) Counts {
	c := make(Counts, 0)
	for k, v := range m {
		c = append(c, &Count{k, v})
	}
	return c
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

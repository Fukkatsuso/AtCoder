package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Work struct {
	A, B int
}

type Works []*Work

func (w Works) Len() int           { return len(w) }
func (w Works) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w Works) Less(i, j int) bool { return (w[i].A - w[i].B) < (w[j].A - w[j].B) }

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, t := nextInt(), nextInt()
	works := make(Works, n)
	sumA, sumB := 0, 0
	for i := range works {
		a, b := nextInt(), nextInt()
		works[i] = &Work{a, b}
		sumA += a
		sumB += b
	}
	sort.Sort(works)

	if sumB > t {
		puts(-1)
		return
	}

	ans := 0
	for i, dt := n-1, 0; i >= 0 && sumA-dt > t; i-- {
		dt += works[i].A - works[i].B
		ans++
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

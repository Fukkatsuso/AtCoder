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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

type Meal struct {
	a, b, c int
}

type Meals []*Meal

func (a Meals) Len() int           { return len(a) }
func (a Meals) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Meals) Less(i, j int) bool { return a[i].c > a[j].c }

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	c := make(Meals, n)
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		c[i] = &Meal{a, b, a + b}
	}
	sort.Sort(c)

	taka, aoki := 0, 0
	for i := range c {
		if i%2 == 0 {
			taka += c[i].a
		} else {
			aoki += c[i].b
		}
	}

	puts(taka - aoki)
}

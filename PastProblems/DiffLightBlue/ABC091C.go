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

type Point struct {
	x, y int
}

type PR []Point

func (a PR) Len() int      { return len(a) }
func (a PR) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PR) Less(i, j int) bool {
	if a[i].y == a[j].y {
		return a[i].x > a[j].x
	}
	return a[i].y > a[j].y
}

type PB []Point

func (a PB) Len() int      { return len(a) }
func (a PB) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PB) Less(i, j int) bool {
	if a[i].x == a[j].x {
		return a[i].y < a[j].y
	}
	return a[i].x < a[j].x
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	pr, pb := make(PR, n), make(PB, n)
	for i := 0; i < n; i++ {
		pr[i] = Point{nextInt(), nextInt()}
	}
	for i := 0; i < n; i++ {
		pb[i] = Point{nextInt(), nextInt()}
	}
	sort.Sort(pr)
	sort.Sort(pb)

	ans := 0
	used := make([]bool, n)
	for i := 0; i < n; i++ {
		// pbを基準にする
		for j := 0; j < n; j++ {
			if used[j] {
				continue
			}
			if pb[i].x > pr[j].x && pb[i].y > pr[j].y {
				ans++
				used[j] = true
				break
			}
		}
	}

	puts(ans)
}

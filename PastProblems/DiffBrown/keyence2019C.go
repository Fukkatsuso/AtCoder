// シミュレーション不要だけど一応AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Elem struct {
	A, B, C int
}

type Elems []*Elem

func (e Elems) Len() int {
	return len(e)
}

func (e Elems) Less(i, j int) bool {
	return (e[i].A - e[i].B) < (e[j].A - e[j].B)
}

func (e Elems) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a, b := nextInts(n), nextInts(n)
	elems := make(Elems, n)
	sumA, sumB := 0, 0
	for i := 0; i < n; i++ {
		elems[i] = &Elem{a[i], b[i], a[i]}
		sumA, sumB = sumA+a[i], sumB+b[i]
	}
	sort.Sort(elems)

	if sumA < sumB {
		puts(-1)
		return
	}

	// Cを調整
	for inc, dec := 0, n-1; dec >= 0 && inc < n && inc <= dec; {
		dinc, ddec := elems[inc].B-elems[inc].C, elems[dec].C-elems[dec].B
		if dinc <= 0 {
			break
		}
		if dinc == ddec {
			elems[inc].C += dinc
			elems[dec].C -= ddec
			dec--
			inc++
		} else if dinc < ddec {
			elems[inc].C += dinc
			elems[dec].C -= dinc
			inc++
		} else {
			elems[inc].C += ddec
			elems[dec].C -= ddec
			dec--
		}
	}

	ans := 0
	for _, e := range elems {
		if e.A != e.C {
			ans++
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

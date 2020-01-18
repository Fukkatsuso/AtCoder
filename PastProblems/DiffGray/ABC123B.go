package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	t := make(MyInts, 5)
	for i := range t {
		t[i] = nextInt()
	}
	sort.Sort(t)

	a := filterA(t)
	b := filterB(t)
	ans := sum(a)
	for i := len(b) - 1; i > 0; i-- {
		ans += b[i] + (10 - b[i]%10)
	}
	if len(b) > 0 {
		ans += b[0]
	}
	puts(ans)
	wt.Flush()
}

type MyInts []int

func (m MyInts) Len() int {
	return len(m)
}

func (m MyInts) Less(i, j int) bool {
	return m[i]%10 < m[j]%10
}

func (m MyInts) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func filterA(m MyInts) []int {
	ret := make([]int, 0)
	for i := range m {
		if m[i]%10 == 0 {
			ret = append(ret, m[i])
		}
	}
	return ret
}

func filterB(m MyInts) []int {
	ret := make([]int, 0)
	for i := range m {
		if m[i]%10 != 0 {
			ret = append(ret, m[i])
		}
	}
	return ret
}

var (
	sc  = bufio.NewScanner(os.Stdin)
	rdr = bufio.NewReaderSize(os.Stdin, 1000000)
	wt  = bufio.NewWriter(os.Stdout)
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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func sum(a []int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

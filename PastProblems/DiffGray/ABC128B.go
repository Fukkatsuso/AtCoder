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
	n := nextInt()
	rs := make(Rests, n)
	for i := range rs {
		rs[i] = &Rest{s: next(), p: nextInt(), id: i + 1}
	}
	sort.Sort(rs)
	for i := range rs {
		puts(rs[i].id)
	}
	wt.Flush()
}

type Rest struct {
	s  string
	p  int
	id int
}

type Rests []*Rest

func (rs Rests) Len() int {
	return len(rs)
}

func (rs Rests) Less(i, j int) bool {
	if rs[i].s == rs[j].s {
		return rs[i].p > rs[j].p
	}
	return rs[i].s < rs[j].s
}

func (rs Rests) Swap(i, j int) {
	rs[i], rs[j] = rs[j], rs[i]
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

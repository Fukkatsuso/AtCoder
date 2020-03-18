package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Dics []*Dic

type Dic struct {
	word  string
	index int
}

func (d Dics) Len() int {
	return len(d)
}

func (d Dics) Less(i, j int) bool {
	return d[i].word < d[j].word
}

func (d Dics) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	s := make([]string, n)
	t := make(Dics, n)
	for i := range s {
		s[i] = next()
		t[i] = &Dic{reverseString(s[i]), i}
	}
	sort.Sort(t)

	for i := range t {
		puts(s[t[i].index])
	}
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

func reverseString(s string) string {
	n := len(s)
	bytes := []byte(s)
	for i := 0; i < n/2; i++ {
		bytes[i], bytes[n-1-i] = bytes[n-1-i], bytes[i]
	}
	return string(bytes)
}

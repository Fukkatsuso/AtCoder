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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

// 1-indexed
type BIT [2][]int

func NewBIT(n int) *BIT {
	var b BIT
	for i := range b {
		b[i] = make([]int, n)
	}
	return &b
}

// [l,r)にxを加算
func (b *BIT) add(l, r, x int) {
	add_sub := func(p, idx, x int) {
		for idx < len(b[p]) {
			b[p][idx] += x
			idx += idx & (-idx)
		}
	}
	add_sub(0, l, -x*(l-1))
	add_sub(0, r, x*(r-1))
	add_sub(1, l, x)
	add_sub(1, r, -x)
}

// [l,r)の和
func (b *BIT) sum(l, r int) int {
	sum_sub := func(p, idx int) int {
		s := 0
		for idx > 0 {
			s += b[p][idx]
			idx -= idx & (-idx)
		}
		return s
	}
	return sum_sub(0, r-1) + sum_sub(1, r-1)*(r-1) - (sum_sub(0, l-1) + sum_sub(1, l-1)*(l-1))
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	s := []byte(next())
	// bits[i]: 文字iについてのBIT
	bits := make([]*BIT, 26)
	for i := range bits {
		bits[i] = NewBIT(n + 1)
	}
	for i := range s {
		bits[int(s[i]-'a')].add(i+1, i+2, 1)
	}

	q := nextInt()
	for ; q > 0; q-- {
		t := nextInt()
		if t == 1 {
			i, c := nextInt()-1, next()[0]
			if s[i] == c {
				continue
			}
			bits[int(s[i]-'a')].add(i+1, i+2, -1)
			bits[int(c-'a')].add(i+1, i+2, 1)
			s[i] = c
		} else {
			l, r := nextInt(), nextInt()
			res := 0
			for i := 0; i < 26; i++ {
				if bits[i].sum(l, r+1) > 0 {
					res++
				}
			}
			puts(res)
		}
	}
}

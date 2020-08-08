// 解説AC

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

type Room struct {
	s, t int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()

	rooms := make([]Room, m)
	imos := make([]int, n+2)
	for i := 0; i < m; i++ {
		s, t := nextInt(), nextInt()
		rooms[i] = Room{s, t}
		imos[s]++
		imos[t+1]--
	}
	for i := 1; i < n+2; i++ {
		imos[i] += imos[i-1]
	}

	// 担当者が1名のみの部屋の個数の累積和
	covered := make([]int, n+1)
	for i := range covered {
		if imos[i] == 1 {
			covered[i] = 1
		}
		if i > 0 {
			covered[i] += covered[i-1]
		}
	}

	// 区間[s,t]に担当者が1名のみの部屋が含まれていたらNG
	ids := make([]int, 0)
	for i := 0; i < m; i++ {
		s, t := rooms[i].s, rooms[i].t
		if covered[t]-covered[s-1] == 0 {
			ids = append(ids, i+1)
		}
	}

	puts(len(ids))
	for i := range ids {
		puts(ids[i])
	}
}

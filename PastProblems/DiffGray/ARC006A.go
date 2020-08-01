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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	e := nextInts(6)
	b := nextInt()
	l := nextInts(6)

	mp := map[int]int{}
	bonus := false
	for i := range e {
		mp[e[i]]++
		mp[l[i]]++
		bonus = bonus || (l[i] == b)
	}

	cnt := 0
	for _, v := range mp {
		if v == 2 {
			cnt++
		}
	}
	ans := 0
	switch cnt {
	case 6:
		ans = 1
	case 5:
		if bonus {
			ans = 2
		} else {
			ans = 3
		}
	case 4:
		ans = 4
	case 3:
		ans = 5
	}
	puts(ans)
}

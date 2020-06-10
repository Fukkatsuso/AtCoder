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

func primeFactor(n int) map[int]int {
	m := map[int]int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	if n != 1 {
		m[n] = 1
	}
	return m
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	pfs := map[int]int{}
	for i := 2; i <= n; i++ {
		pf := primeFactor(i)
		for k, v := range pf {
			pfs[k] += v
		}
	}

	num := func(m int) int {
		res := 0
		for _, v := range pfs {
			if v >= m-1 {
				res++
			}
		}
		return res
	}

	ans := num(75) + num(25)*(num(3)-1) + num(15)*(num(5)-1) + num(5)*(num(5)-1)*(num(3)-2)/2
	puts(ans)
}

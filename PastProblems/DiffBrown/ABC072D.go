// 解説AC
// p[n-1]を探索範囲に含めずWA

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	p := make([]int, n)
	for i := range p {
		p[i] = nextInt() - 1
	}

	ans := 0
	for i := 0; i < n; i++ {
		if p[i] == i {
			if i < n-1 {
				p[i], p[i+1] = p[i+1], p[i]
			}
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

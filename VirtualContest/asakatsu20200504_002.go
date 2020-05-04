// ARC014B

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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	w := make([]string, n)
	for i := range w {
		w[i] = next()
	}

	used := map[string]bool{}
	for i := range w {
		if i%2 == 0 {
			if used[w[i]] || (i > 0 && w[i][0] != w[i-1][len(w[i-1])-1]) {
				puts("LOSE")
				return
			} else {
				used[w[i]] = true
			}
		} else {
			if used[w[i]] || (i > 0 && w[i][0] != w[i-1][len(w[i-1])-1]) {
				puts("WIN")
				return
			} else {
				used[w[i]] = true
			}
		}
	}

	puts("DRAW")
}

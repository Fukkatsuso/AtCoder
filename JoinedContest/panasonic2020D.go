// WA
// dfsの伸ばし方が惜しい

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var n int

func dfs(s []byte, x, kind int) {
	if x == n {
		putf("%s\n", string(s))
		return
	}
	for i := 0; i < kind; i++ {
		s[x] = byte('a' + i)
		dfs(s, x+1, kind+1)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n = nextInt()

	dfs(make([]byte, n), 0, 1)
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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

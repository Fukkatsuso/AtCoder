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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	edges := make([][2]int, 0)
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if n%2 == 1 && i+j != n {
				edges = append(edges, [2]int{i, j})
			} else if n%2 == 0 && i+j != n+1 {
				edges = append(edges, [2]int{i, j})
			}
		}
	}

	puts(len(edges))
	for _, edge := range edges {
		puts(edge[0], edge[1])
	}
}

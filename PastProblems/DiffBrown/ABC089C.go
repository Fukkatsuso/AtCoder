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
	name := map[byte]int{}
	for i := 0; i < n; i++ {
		s := next()
		name[s[0]]++
	}

	initials := "MARCH"
	ans := 0
	for i := 0; i < len(initials)-2; i++ {
		for j := i + 1; j < len(initials)-1; j++ {
			for k := j + 1; k < len(initials); k++ {
				ans += name[initials[i]] * name[initials[j]] * name[initials[k]]
			}
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

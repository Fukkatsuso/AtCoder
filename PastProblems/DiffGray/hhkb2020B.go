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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w := getInt(), getInt()
	s := make([]string, h)
	for i := range s {
		s[i] = gets()
	}

	ans := 0
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if i < h-1 && s[i][j] == '.' && s[i+1][j] == '.' {
				ans++
			}
			if j < w-1 && s[i][j] == '.' && s[i][j+1] == '.' {
				ans++
			}
		}
	}

	puts(ans)
}

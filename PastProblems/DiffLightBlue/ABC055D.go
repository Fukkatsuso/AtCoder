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

func ok(n int, s string, t []byte) bool {
	for i := 1; i < n-1; i++ {
		if (t[i] == 'S' && s[i] == 'o') || (t[i] == 'W' && s[i] == 'x') {
			t[i+1] = t[i-1]
		} else {
			if t[i-1] == 'S' {
				t[i+1] = 'W'
			} else {
				t[i+1] = 'S'
			}
		}
	}
	// n-1の両隣
	if (t[n-1] == 'S' && s[n-1] == 'o') || (t[n-1] == 'W' && s[n-1] == 'x') {
		if t[n-2] != t[0] {
			return false
		}
	} else {
		if t[n-2] == t[0] {
			return false
		}
	}
	// 0の両隣
	if (t[0] == 'S' && s[0] == 'o') || (t[0] == 'W' && s[0] == 'x') {
		if t[n-1] != t[1] {
			return false
		}
	} else {
		if t[n-1] == t[1] {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, s := nextInt(), next()

	t := make([]byte, n)
	sw := []byte{'S', 'W'}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			t[0], t[1] = sw[i], sw[j]
			if ok(n, s, t) {
				puts(string(t))
				return
			}
		}
	}

	puts(-1)
}

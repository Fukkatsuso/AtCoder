package main

import (
	"bufio"
	"fmt"
	"os"
)

func atcoder(c byte) bool {
	switch c {
	case 'a', 't', 'c', 'o', 'd', 'e', 'r':
		return true
	default:
		return false
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s, t := next(), next()

	ok := true
	for i := range s {
		if s[i] == t[i] {
			continue
		}
		if s[i] == '@' && !atcoder(t[i]) {
			ok = false
		} else if !atcoder(s[i]) && t[i] == '@' {
			ok = false
		} else if s[i] != '@' && t[i] != '@' {
			ok = false
		}

	}

	if ok {
		puts("You can win")
	} else {
		puts("You will lose")
	}
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := next()
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if i == 0 {
				putf(",")
			} else if s[i-1] != ' ' {
				putf(",")
			}
		} else {
			putf("%c", s[i])
		}
	}
	putf("\n")
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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

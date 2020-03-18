package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func match(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		if t[i] != s[i] && t[i] != '*' {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanLines)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := strings.Split(next(), " ")
	n := nextInt()
	t := make([]string, n)
	for i := range t {
		t[i] = next()
	}

	for i := range s {
		filtered := false
		for j := range t {
			if match(s[i], t[j]) {
				filtered = true
				break
			}
		}
		if filtered {
			s[i] = strings.Repeat("*", len(s[i]))
		}
		if i < len(s)-1 {
			putf("%s ", s[i])
		} else {
			putf("%s\n", s[i])
		}
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

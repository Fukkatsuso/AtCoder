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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func conv(b byte) string {
	switch b {
	case 'z', 'r', 'Z', 'R':
		return "0"
	case 'b', 'c', 'B', 'C':
		return "1"
	case 'd', 'w', 'D', 'W':
		return "2"
	case 't', 'j', 'T', 'J':
		return "3"
	case 'f', 'q', 'F', 'Q':
		return "4"
	case 'l', 'v', 'L', 'V':
		return "5"
	case 's', 'x', 'S', 'X':
		return "6"
	case 'p', 'm', 'P', 'M':
		return "7"
	case 'h', 'k', 'H', 'K':
		return "8"
	case 'n', 'g', 'N', 'G':
		return "9"
	}
	return ""
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	ans := make([]string, 0)
	for i := 0; i < n; i++ {
		w := next()
		s := ""
		for j := range w {
			s += conv(w[j])
		}

		if len(s) > 0 {
			ans = append(ans, s)
		}
	}

	if len(ans) == 0 {
		putf("\n")
		return
	}

	for i := range ans {
		if i < len(ans)-1 {
			putf("%s ", ans[i])
		} else {
			putf("%s\n", ans[i])
		}
	}
}

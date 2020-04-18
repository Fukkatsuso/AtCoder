package main

import (
	"bufio"
	"fmt"
	"os"
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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := next()
	n := len(s)

	rsf := map[byte]map[string]bool{
		'S': map[string]bool{},
		'H': map[string]bool{},
		'D': map[string]bool{},
		'C': map[string]bool{},
	}
	allGet := func(mark byte) bool {
		return rsf[mark]["10"] && rsf[mark]["J"] && rsf[mark]["Q"] && rsf[mark]["K"] && rsf[mark]["A"]
	}

	var ansMark byte
	var lastIndex int
	for i := 0; i < n; {
		mark := s[i]
		var num string
		if s[i+1] == '1' {
			num = "10"
			i += 3
		} else {
			num = string(s[i+1])
			i += 2
		}
		rsf[mark][num] = true
		if allGet(mark) {
			ansMark = mark
			lastIndex = i
			break
		}
	}

	ans := ""
	got := map[string]bool{}
	for i, nextI := 0, 0; i < lastIndex; i = nextI {
		mark := s[i]
		var num string
		if s[i+1] == '1' {
			num = "10"
			nextI = i + 3
		} else {
			num = string(s[i+1])
			nextI = i + 2
		}
		if mark != ansMark {
			ans += string(mark) + num
			continue
		}
		switch num {
		case "10", "J", "Q", "K", "A":
			if got[num] {
				ans += string(mark) + num
			} else {
				got[num] = true
			}
		default:
			ans += string(mark) + num
		}
	}

	if len(ans) == 0 {
		puts(0)
	} else {
		puts(ans)
	}
}

// sc.Split()とnextMega()の相性が悪いらしい

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	s := nextMega()

	rank := 0
	sum := 0
	for i := 0; i < n; i++ {
		ok := false
		if s[i] == 'a' && sum < a+b {
			sum++
			ok = true
		} else if s[i] == 'b' {
			rank++
			if sum < a+b && rank <= b {
				sum++
				ok = true
			}
		}
		if ok {
			puts("Yes")
		} else {
			puts("No")
		}
	}
	wt.Flush()
}

var (
	sc  = bufio.NewScanner(os.Stdin)
	rdr = bufio.NewReaderSize(os.Stdin, 1000000)
	wt  = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextMega() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, _ := rdr.ReadLine()
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

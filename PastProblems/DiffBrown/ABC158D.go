package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := next()
	l := list.New()
	for _, c := range s {
		l.PushBack(string(c))
	}
	q := nextInt()

	rev := false

	for i := 0; i < q; i++ {
		t := nextInt()
		switch t {
		case 1:
			rev = !rev
		case 2:
			f, c := nextInt(), next()
			if (!rev && f == 1) || (rev && f == 2) {
				l.PushFront(c)
			} else {
				l.PushBack(c)
			}
		}
	}

	if !rev {
		for e := l.Front(); e != nil; e = e.Next() {
			putf("%s", e.Value.(string))
		}
	} else {
		for e := l.Back(); e != nil; e = e.Prev() {
			putf("%s", e.Value.(string))
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

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

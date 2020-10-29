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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

const (
	dict    = "dict"
	set     = "set"
	integer = "int"
	expr    = "expr"
	colon   = ":"
	comma   = ","
	open    = "{"
	close   = "}"
	empty   = "empty"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := next()

	q := make([]string, 0)
	push := func(t string) {
		q = append(q, t)
	}
	top := func() string {
		if len(q) == 0 {
			return empty
		}
		return q[len(q)-1]
	}
	pop := func() string {
		t := top()
		if len(q) > 0 {
			q = q[:len(q)-1]
		}
		return t
	}

	for i := 0; i < len(s); i++ {
		b := s[i]
		switch b {
		case '{':
			push(open)
		case '}':
			// open が出てくるまで取り出す
			d := (top() == open)
			for {
				t := pop()
				d = d || (t == colon)
				if t == open {
					break
				}
			}
			if d {
				push(dict)
			} else {
				push(set)
			}
		case ':':
			push(colon)
		case ',':
			push(comma)
		default:
			if top() != integer {
				push(integer)
			}
		}
	}

	puts(top())
}

// DICT = "{" , EXPR , ":" , EXPR , { "," , EXPR , ":" , EXPR } , "}" | "{}" ;
// SET  = "{" , EXPR , { "," , EXPR } , "}" ;
// EXPR = DICT | SET | INTEGER ;

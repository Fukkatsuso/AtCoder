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

func nextXY(n, x, y int, w string) (int, int, string) {
	switch w {
	case "R":
		x++
		if x == n {
			return n - 2, y, "L"
		}
	case "L":
		x--
		if x == -1 {
			return 1, y, "R"
		}
	case "U":
		y--
		if y == -1 {
			return x, 1, "D"
		}
	case "D":
		y++
		if y == n {
			return x, n - 2, "U"
		}
	case "RU":
		x, y = x+1, y-1
		if x == n && y == -1 {
			return n - 2, 1, "LD"
		} else if x == n {
			return n - 2, y, "LU"
		} else if y == -1 {
			return x, 1, "RD"
		}
	case "RD":
		x, y = x+1, y+1
		if x == n && y == n {
			return n - 2, n - 2, "LU"
		} else if x == n {
			return n - 2, y, "LD"
		} else if y == n {
			return x, n - 2, "RU"
		}
	case "LU":
		x, y = x-1, y-1
		if x == -1 && y == -1 {
			return 1, 1, "RD"
		} else if x == -1 {
			return 1, y, "RU"
		} else if y == -1 {
			return x, 1, "LD"
		}
	case "LD":
		x, y = x-1, y+1
		if x == -1 && y == n {
			return 1, n - 2, "RU"
		} else if x == -1 {
			return 1, y, "RD"
		} else if y == n {
			return x, n - 2, "LU"
		}
	}
	return x, y, w
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	x, y, w := nextInt()-1, nextInt()-1, next()
	c := make([]string, 9)
	for i := range c {
		c[i] = next()
	}

	ans := make([]byte, 4)
	for cnt := 0; cnt < 4; cnt++ {
		ans[cnt] = c[y][x]
		x, y, w = nextXY(9, x, y, w)
	}

	putf("%s\n", ans)
}

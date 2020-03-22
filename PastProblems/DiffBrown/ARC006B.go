package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanLines)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	var n, l int
	fmt.Scan(&n, &l)
	lines := make([]string, l)
	for i := range lines {
		lines[i] = next()
	}
	y := next()

	var ans int
	for i := 0; i < n*2; i += 2 {
		if y[i] == 'o' {
			ans = i / 2
			break
		}
	}

	// move: -1, 0, +1
	for i, move := l-1, 2; i >= 0; i, move = i-1, 2 {
		line := lines[i]
		for move != 0 {
			if left := ans*2 - 1; left > 0 && line[left] == '-' && move != 1 {
				ans--
				move = -1
			} else if right := ans*2 + 1; right < n*2-1 && line[right] == '-' && move != -1 {
				ans++
				move = 1
			} else {
				move = 0
			}
		}
	}

	puts(ans + 1)
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

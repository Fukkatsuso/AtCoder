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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	s := next()

	rgb := "RGB"
	nxt := make([][3]int, n)
	nxt[n-1][0], nxt[n-1][1], nxt[n-1][2] = n, n, n
	for i := n - 2; i >= 0; i-- {
		for j := range rgb {
			if s[i+1] == rgb[j] {
				nxt[i][j] = i + 1
			} else {
				nxt[i][j] = nxt[i+1][j]
			}
		}
	}

	b := []byte{s[0]}
	for i := 1; i < n; i++ {
		l := len(b)
		if l == 0 {
			b = append(b, s[i])
			continue
		} else if s[i] == b[0] {
			b = b[1:]
			continue
		} else if s[i] == b[l-1] {
			b = b[:l-1]
			continue
		}
		for j := range rgb {
			if s[i] != rgb[j] {
				continue
			}
			if b[0] == rgb[(j+1)%3] {
				if nxt[i][(j+1)%3] > nxt[i][(j+2)%3] {
					b = append([]byte{s[i]}, b...)
				} else {
					b = append(b, s[i])
				}
			} else {
				if nxt[i][(j+1)%3] < nxt[i][(j+2)%3] {
					b = append([]byte{s[i]}, b...)
				} else {
					b = append(b, s[i])
				}
			}
		}
	}

	puts(len(b))
}

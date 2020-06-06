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

func number(s1, s2, s3, s4, s5 string) int {
	seg1 := ".###..#..###.###.#.#.###.###.###.###.###."
	seg2 := ".#.#.##....#...#.#.#.#...#.....#.#.#.#.#."
	seg3 := ".#.#..#..###.###.###.###.###...#.###.###."
	seg4 := ".#.#..#..#.....#...#...#.#.#...#.#.#...#."
	seg5 := ".###.###.###.###...#.###.###...#.###.###."
	var x int
	for i := 0; i < 10; i++ {
		if s1 != seg1[4*i:4*(i+1)] {
			continue
		}
		if s2 != seg2[4*i:4*(i+1)] {
			continue
		}
		if s3 != seg3[4*i:4*(i+1)] {
			continue
		}
		if s4 != seg4[4*i:4*(i+1)] {
			continue
		}
		if s5 != seg5[4*i:4*(i+1)] {
			continue
		}
		x = i
		break
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	s := make([]string, 5)
	for i := range s {
		s[i] = next()
	}

	num := make([]int, n)
	for i := range num {
		num[i] = number(s[0][4*i:4*(i+1)], s[1][4*i:4*(i+1)], s[2][4*i:4*(i+1)], s[3][4*i:4*(i+1)], s[4][4*i:4*(i+1)])
	}

	for i := range num {
		putf("%d", num[i])
	}
	puts()
}

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

func nextInt2() (int, int) {
	return nextInt(), nextInt()
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w := nextInt2()
	a := make([][]int, h)
	for i := range a {
		a[i] = nextInts(w)
	}

	ansN := 0
	op := make([][4]int, 0)

	for i := 0; i < h; i++ {
		for j := (i % 2) * (w - 1); 0 <= j && j < w; j += 1 - 2*(i%2) {
			if a[i][j]%2 == 0 {
				continue
			}
			var nextX, nextY int
			if (j == w-1 && i%2 == 0) || (j == 0 && i%2 == 1) {
				nextX, nextY = j, i+1
			} else {
				nextX, nextY = j+1-2*(i%2), i
			}
			if nextY >= h {
				continue
			}
			a[nextY][nextX]++
			op = append(op, [4]int{i + 1, j + 1, nextY + 1, nextX + 1})
			ansN++
		}
	}

	puts(ansN)
	for i := range op {
		puts(op[i][0], op[i][1], op[i][2], op[i][3])
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(a ...int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	a := make([][]int, 3)
	for i := range a {
		a[i] = nextInts(3)
	}
	n := nextInt()

	for i := 0; i < n; i++ {
		b := nextInt()
		for j := 0; j < 9; j++ {
			if b == a[j/3][j%3] {
				a[j/3][j%3] = 0
			}
		}
	}

	bingo := false
	for i := 0; i < 3; i++ {
		if sum(a[i][0], a[i][1], a[i][2]) == 0 {
			bingo = true
		}
	}
	for i := 0; i < 3; i++ {
		if sum(a[0][i], a[1][i], a[2][i]) == 0 {
			bingo = true
		}
	}
	if sum(a[0][0], a[1][1], a[2][2]) == 0 || sum(a[0][2], a[1][1], a[2][0]) == 0 {
		bingo = true
	}

	if bingo {
		puts("Yes")
	} else {
		puts("No")
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
	i, _ := strconv.Atoi(next())
	return i
}

func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func nextFloat64s(n int) []float64 {
	slice := make([]float64, n)
	for i := 0; i < n; i++ {
		slice[i] = nextFloat64()
	}
	return slice
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

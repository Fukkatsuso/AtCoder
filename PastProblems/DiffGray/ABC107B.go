package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	h, w := nextInt(), nextInt()
	a := nextStrings(h)

	whiteh, whitew := make([]bool, h), make([]bool, w)
	for i := 0; i < h; i++ {
		white := true
		for j := 0; j < w; j++ {
			if a[i][j] == '#' {
				white = false
			}
		}
		whiteh[i] = white
	}
	for j := 0; j < w; j++ {
		white := true
		for i := 0; i < h; i++ {
			if a[i][j] == '#' {
				white = false
			}
		}
		whitew[j] = white
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if !whiteh[i] && !whitew[j] {
				putf("%c", a[i][j])
			}
		}
		if !whiteh[i] {
			putf("\n")
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

func nextStrings(n int) []string {
	slice := make([]string, n)
	for i := 0; i < n; i++ {
		slice[i] = next()
	}
	return slice
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n, s := nextInt(), next()
	up, down := make([][]bool, n), make([][]bool, n)
	for i := 0; i < n; i++ {
		up[i], down[i] = make([]bool, 26), make([]bool, 26)
	}

	up[0][s[0]-'a'] = true
	for i := 1; i < n; i++ {
		for j := 0; j < 26; j++ {
			up[i][j] = up[i-1][j]
		}
		up[i][s[i]-'a'] = true
	}
	down[n-1][s[n-1]-'a'] = true
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			down[i][j] = down[i+1][j]
		}
		down[i][s[i]-'a'] = true
	}

	ans := 0
	for i := 1; i < n-1; i++ {
		cnt := 0
		for j := 0; j < 26; j++ {
			if up[i][j] == true && down[i+1][j] == true {
				cnt++
			}
		}
		if cnt > ans {
			ans = cnt
		}
	}
	puts(ans)
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

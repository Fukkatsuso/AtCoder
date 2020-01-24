package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Drink struct {
	price, num int
}

type Drinks []*Drink

func (d Drinks) Len() int {
	return len(d)
}

func (d Drinks) Less(i, j int) bool {
	return d[i].price < d[j].price
}

func (d Drinks) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	d := make(Drinks, n)
	for i := range d {
		d[i] = &Drink{price: nextInt(), num: nextInt()}
	}
	sort.Sort(d)

	ans := 0
	for i := 0; m > 0; i++ {
		if d[i].num >= m {
			ans += m * d[i].price
			m = 0
		} else {
			ans += d[i].num * d[i].price
			m -= d[i].num
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

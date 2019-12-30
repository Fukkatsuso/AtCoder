package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
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

func main() {
	sc.Split(bufio.ScanWords)
	km := nextFloat64() / 1000.0
	ans := ""
	if km < 0.1 {
		ans = "00"
	} else if km <= 5 {
		ans = fmt.Sprintf("%02d", int(km*10))
	} else if km <= 30 {
		ans = fmt.Sprintf("%d", int(km+50))
	} else if km <= 70 {
		ans = fmt.Sprintf("%d", int((km-30)/5+80))
	} else {
		ans = "89"
	}
	fmt.Println(ans)
}

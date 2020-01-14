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

func min(nums []int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a, b := nextInt(), nextInt()
	cnt := make([]int, 3)
	for i := 0; i < n; i++ {
		p := nextInt()
		if p <= a {
			cnt[0]++
		} else if p <= b {
			cnt[1]++
		} else {
			cnt[2]++
		}
	}
	fmt.Println(min(cnt))
}

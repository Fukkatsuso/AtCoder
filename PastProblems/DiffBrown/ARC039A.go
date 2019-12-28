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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func pow(x, y int) int {
	ret := 1
	for i := 0; i < y; i++ {
		ret *= x
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	a, b := next(), next()
	atoi, _ := strconv.Atoi(a)
	btoi, _ := strconv.Atoi(b)
	max := atoi - btoi
	dif := atoi - btoi

	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if i == 0 && j == 0 {
				continue
			}
			// aのi桁目をjに置換したときのa-b
			preAi := int(a[i] - '0')
			c := dif + (j-preAi)*pow(10, 2-i)
			if c > max {
				max = c
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if i == 0 && j == 0 {
				continue
			}
			// bのi桁目をjに置換したときのa-b
			preBi := int(b[i] - '0')
			c := dif - (j-preBi)*pow(10, 2-i)
			if c > max {
				max = c
			}
		}
	}

	fmt.Println(max)
}

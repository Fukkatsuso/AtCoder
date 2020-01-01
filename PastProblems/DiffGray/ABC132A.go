// 汚すぎる
// 等しい・等しくないの組3通りを試せばよい

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
	s := next()
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i == j {
				continue
			}
			if s[i] == s[j] {
				for k := 0; k < 4; k++ {
					if k == i || k == j {
						continue
					}
					l := 6 - i - j - k
					if s[i] != s[k] && s[k] == s[l] {
						fmt.Println("Yes")
						return
					}
				}
			}
		}
	}
	fmt.Println("No")
}

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
	n := nextInt()

	// 4の倍数の個数, 奇数の個数
	x, y := 0, 0
	for i := 0; i < n; i++ {
		a := nextInt()
		if a%4 == 0 {
			x++
		} else if a%2 == 1 {
			y++
		}
	}

	// 第1の条件を満たすなら必然的に第2の条件も満たすから、第1条件は不要かも
	if x >= n/2 {
		fmt.Println("Yes")
	} else if x >= y {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
